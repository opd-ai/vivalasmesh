// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// handleKey handles keyboard input.
func (m *Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case matches(msg, m.Keys.Up):
		return m, m.moveCursor(0, -1)
	case matches(msg, m.Keys.Down):
		return m, m.moveCursor(0, 1)
	case matches(msg, m.Keys.Left):
		return m, m.moveCursor(-1, 0)
	case matches(msg, m.Keys.Right):
		return m, m.moveCursor(1, 0)
	case matches(msg, m.Keys.SmartKey):
		return m, m.smartAction()
	case matches(msg, m.Keys.Cancel):
		return m, m.cancelAction()
	case matches(msg, m.Keys.Command):
		m.CurrentMode = ModeCommandPalette
		return m, nil
	case matches(msg, m.Keys.Tab):
		m.cycleMode()
		return m, nil
	case matches(msg, m.Keys.Help):
		m.CurrentMode = ModeHelpOverlay
		return m, nil
	}
	return m, nil
}

// handleMouse handles mouse input.
func (m *Model) handleMouse(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	if !m.MouseEnabled {
		return m, nil
	}

	switch msg.Type {
	case tea.MouseLeft:
		return m, m.handleLeftClick(msg.X, msg.Y)
	case tea.MouseRight:
		return m, m.handleRightClick(msg.X, msg.Y)
	case tea.MouseMotion:
		return m, m.handleMouseMove(msg.X, msg.Y)
	}
	return m, nil
}

// moveCursor moves the cursor/selection.
func (m *Model) moveCursor(dx, dy int) tea.Cmd {
	// In real implementation, this would move the player or cursor
	return nil
}

// smartAction triggers the contextual smart key action.
func (m *Model) smartAction() tea.Cmd {
	// In real implementation, this would resolve the highest-priority contextual action
	return nil
}

// cancelAction cancels the current action/prompt.
func (m *Model) cancelAction() tea.Cmd {
	switch m.CurrentMode {
	case ModeCommandPalette, ModeRadialMenu, ModeHelpOverlay, ModeOnboarding:
		m.CurrentMode = ModeMapView
	}
	return nil
}

// cycleMode cycles through focus modes.
func (m *Model) cycleMode() {
	modes := []Mode{ModeMapView, ModeInventory, ModeSpectrumAnalyzer}
	for i, mode := range modes {
		if mode == m.CurrentMode {
			m.CurrentMode = modes[(i+1)%len(modes)]
			break
		}
	}
}

// handleLeftClick handles left mouse click.
func (m *Model) handleLeftClick(x, y int) tea.Cmd {
	// In real implementation: point-and-click A* pathfinding
	return nil
}

// handleRightClick handles right mouse click.
func (m *Model) handleRightClick(x, y int) tea.Cmd {
	// In real implementation: context radial menu
	m.CurrentMode = ModeRadialMenu
	return nil
}

// handleMouseMove handles mouse movement.
func (m *Model) handleMouseMove(x, y int) tea.Cmd {
	// In real implementation: hover effects, tooltips
	return nil
}

func matches(msg tea.KeyMsg, binding key.Binding) bool {
	for _, k := range binding.Keys() {
		if msg.String() == k {
			return true
		}
	}
	return false
}
