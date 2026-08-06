// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the main Bubble Tea model.
type Model struct {
	CurrentMode  Mode
	Width        int
	Height       int
	Keys         KeyMap
	MouseEnabled bool
}

// NewModel creates a new TUI model.
func NewModel() *Model {
	return &Model{
		CurrentMode:  ModeMapView,
		Keys:         DefaultKeyMap(),
		MouseEnabled: true,
	}
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	// Enable mouse tracking (SGR 1006 mode)
	return tea.EnableMouseAllMotion
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
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
	// In real implementation, this would render the sub-pixel canvas
	return "Viva Las Mesh - " + m.CurrentMode.String()
}
