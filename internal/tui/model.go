// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	render "github.com/opd-ai/vivalasmesh/internal/render"
)

// Model is the main Bubble Tea model.
type Model struct {
	CurrentMode  Mode
	Width        int
	Height       int
	Canvas       *render.Canvas
	Keys         KeyMap
	MouseEnabled bool
}

// NewModel creates a new TUI model.
func NewModel() *Model {
	canvas, _ := render.NewCanvas(80, 24)
	return &Model{
		CurrentMode:  ModeMapView,
		Width:        80,
		Height:       24,
		Canvas:       canvas,
		Keys:         DefaultKeyMap(),
		MouseEnabled: true,
	}
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	// Enable mouse tracking (SGR 1006 mode)
	return tea.EnableMouseAllMotion
}

// handleWindowSizeMsg updates the model's dimensions and resizes the canvas on terminal resize.
func (m *Model) handleWindowSizeMsg(msg tea.WindowSizeMsg) {
	m.Width = msg.Width
	m.Height = msg.Height
	if err := m.Canvas.Resize(msg.Width, msg.Height); err != nil {
		// ignore error
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

	default:
		return m, nil
	}
}

// View implements tea.Model.
func (m *Model) View() string {
	return m.Canvas.Render()
}
