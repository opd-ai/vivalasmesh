// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package main is the entry point for Viva Las Mesh.
package main

import (
	"log"

	"github.com/opd-ai/vivalasmesh/internal/engine"
	"github.com/opd-ai/vivalasmesh/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

// main is the application entry point.
func main() {
	// Create a new game engine.
	eng := engine.NewEngine()
	// Start the engine in a separate goroutine.
	if err := eng.Start(); err != nil {
		log.Fatalf("Failed to start engine: %v", err)
	}
	defer eng.Stop()

	// Create a new TUI model.
	m := tui.NewModel()

	// Create a Bubble Tea program with the TUI model.
	p := tea.NewProgram(m)

	// Start the Bubble Tea event loop. This blocks until the program exits.
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
