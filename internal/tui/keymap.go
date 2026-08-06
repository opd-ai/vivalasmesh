// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap holds key bindings.
type KeyMap struct {
	Up       key.Binding
	Down     key.Binding
	Left     key.Binding
	Right    key.Binding
	SmartKey key.Binding
	Cancel   key.Binding
	Command  key.Binding
	Tab      key.Binding
	Help     key.Binding
}

// DefaultKeyMap returns the default key bindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up:       key.NewBinding(key.WithKeys("up", "w"), key.WithHelp("↑/W", "move up")),
		Down:     key.NewBinding(key.WithKeys("down", "s"), key.WithHelp("↓/S", "move down")),
		Left:     key.NewBinding(key.WithKeys("left", "a"), key.WithHelp("←/A", "move left")),
		Right:    key.NewBinding(key.WithKeys("right", "d"), key.WithHelp("→/D", "move right")),
		SmartKey: key.NewBinding(key.WithKeys("enter", "space"), key.WithHelp("Space/Enter", "smart action")),
		Cancel:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("Esc", "cancel/back")),
		Command:  key.NewBinding(key.WithKeys("ctrl+k", "/"), key.WithHelp("Ctrl+K", "command palette")),
		Tab:      key.NewBinding(key.WithKeys("tab"), key.WithHelp("Tab", "cycle focus")),
		Help:     key.NewBinding(key.WithKeys("f1", "?"), key.WithHelp("F1/?", "help")),
	}
}
