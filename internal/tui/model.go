// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
package tui

import (
	"context"
	"image/color"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	render "github.com/opd-ai/vivalasmesh/internal/render"

	transport "github.com/opd-ai/vivalasmesh/internal/transport"
)

// Model is the main Bubble Tea model.
type Model struct {
	CurrentMode  Mode
	Width        int
	Height       int
	Canvas       *render.Canvas
	Keys         KeyMap
	MouseEnabled bool
	// spectrum holds the normalized amplitude values for each frequency bin.
	spectrum []float64
	// spectrumTicker is used to periodically update the spectrum.
	spectrumTicker *time.Ticker
	// rfReceiver is the full-spectrum RF receiver providing live spectrum data.
	rfReceiver transport.Receiver
}

// NewModel creates a new TUI model.
func NewModel() *Model {
	canvas, _ := render.NewCanvas(80, 24)
	m := &Model{
		CurrentMode:  ModeMapView,
		Width:        80,
		Height:       24,
		Canvas:       canvas,
		Keys:         DefaultKeyMap(),
		MouseEnabled: true,
		spectrum:     make([]float64, 120), // 120 frequency bins (one per column)
	}
	// Initialize spectrum with zeros
	for i := range m.spectrum {
		m.spectrum[i] = 0
	}
	// Create a full-spectrum RF receiver
	m.rfReceiver = transport.NewSimpleReceiver()
	// Tune to a default frequency (e.g., 915.5 MHz for Radio Free Graceland)
	_ = m.rfReceiver.Tune(context.Background(), 915.5e6)
	// Start the receiver
	_ = m.rfReceiver.Start(context.Background())
	// Create a ticker to update the spectrum every 100ms.
	m.spectrumTicker = time.NewTicker(100 * time.Millisecond)
	return m
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	// Enable mouse tracking (SGR 1006 mode) - disabled for now due to TTY issues in test environment
	// return tea.Batch(
	// 	tea.EnableMouseAllMotion,
	// 	m.spectrumUpdateCmd(),
	// )
	return m.spectrumUpdateCmd()
}

// spectrumUpdateCmd returns a command that sends a SpectrumUpdateMsg after the ticker ticks.
func (m *Model) spectrumUpdateCmd() tea.Cmd {
	return func() tea.Msg {
		<-m.spectrumTicker.C
		return SpectrumUpdateMsg{}
	}
}

// SpectrumUpdateMsg is a message that triggers a spectrum update.
type SpectrumUpdateMsg struct{}

// handleWindowSizeMsg updates the model's dimensions and resizes the canvas on terminal resize.
func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.Width = msg.Width
	m.Height = msg.Height
	if err := m.Canvas.Resize(msg.Width, msg.Height); err != nil {
		// ignore error
	}
	// Reset spectrum buffer if width changed
	if m.Width != len(m.spectrum) {
		m.spectrum = make([]float64, m.Width)
		for i := range m.spectrum {
			m.spectrum[i] = 0
		}
		// Reset ticker interval if needed (keeping 100ms)
		if m.spectrumTicker != nil {
			m.spectrumTicker.Stop()
		}
		m.spectrumTicker = time.NewTicker(100 * time.Millisecond)
	}
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.handleWindowSizeMsg(msg)
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)

	case tea.MouseMsg:
		return m.handleMouse(msg)

	case SpectrumUpdateMsg:
		// Generate new spectrum data and redraw
		m.generateSpectrum()
		m.drawSpectrum()
		// Request another update
		return m, m.spectrumUpdateCmd()

	default:
		return m, nil
	}
}

// generateSpectrum populates the spectrum buffer with data from the RF receiver.
// In a real implementation, this would come from the RF transports.
func (m *Model) generateSpectrum() {
	// Get spectrum data from the RF receiver
	spec := m.rfReceiver.GetSpectrum()
	// Copy the spectrum data into the model's spectrum buffer
	// Ensure we don't exceed the buffer size
	copyLen := len(spec)
	if copyLen > len(m.spectrum) {
		copyLen = len(m.spectrum)
	}
	copy(m.spectrum, spec[:copyLen])
	// If the receiver returned fewer bins, zero out the rest
	for i := copyLen; i < len(m.spectrum); i++ {
		m.spectrum[i] = 0
	}
}

// drawSpectrum draws the spectrum analyzer on the canvas in the Z4 layer (HUD).
func (m *Model) drawSpectrum() {
	// Clear the Z4 layer (HUD) before drawing
	if err := m.Canvas.Clear(render.Z4HUD); err != nil {
		// Ignore error
	}
	// We'll draw the spectrum as a bar graph in the top 10 cells of the screen.
	// Each column represents a frequency bin.
	// The height of the bar is proportional to the amplitude.
	const maxHeight = 10 // cells
	// We'll draw from the top down (y=0 is top)
	for x := 0; x < m.Width && x < len(m.spectrum); x++ {
		amplitude := m.spectrum[x]
		if amplitude < 0 {
			amplitude = 0
		}
		if amplitude > 1 {
			amplitude = 1
		}
		// Calculate bar height in cells
		barHeight := int(amplitude * float64(maxHeight))
		// Draw the bar from y=0 to y=barHeight-1 (top down)
		// We'll use a color gradient from blue (low) to red (high)
		var fg, bg color.RGBA
		if amplitude < 0.3 {
			// Blue
			fg = color.RGBA{R: 0, G: 0, B: 255, A: 255}
			bg = color.RGBA{R: 0, G: 0, B: 100, A: 255}
		} else if amplitude < 0.7 {
			// Yellow to red gradient
			r := uint8(255 * (amplitude - 0.3) / 0.4)
			g := uint8(255 * (1 - (amplitude-0.3)/0.4))
			fg = color.RGBA{R: r, G: g, B: 0, A: 255}
			bg = color.RGBA{R: r / 2, G: g / 2, B: 0, A: 255}
		} else {
			// Red
			fg = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			bg = color.RGBA{R: 100, G: 0, B: 0, A: 255}
		}
		// Draw each cell in the bar
		for y := 0; y < barHeight; y++ {
			// Each cell is a sub-pixel: we'll set both top and bottom to the same color for a solid block.
			// In half-block rendering, we need to set the top (foreground) and bottom (background).
			// To get a solid block, we set both to the same color.
			sp := render.SubPixel{Top: fg, Bottom: bg}
			if err := m.Canvas.Set(render.Z4HUD, x, y, sp); err != nil {
				// Ignore error
			}
		}
	}
}

// View implements tea.Model.
func (m *Model) View() string {
	return m.Canvas.Render()
}
