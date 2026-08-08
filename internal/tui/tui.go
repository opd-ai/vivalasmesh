// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package tui provides the Bubble Tea TUI for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Bubble Tea TUI, input handling, fuzzy palette.
//
// Implements:
//   - Main Bubble Tea model with pulse integration
//   - Input handling (WASD, arrows, Space/Enter, Ctrl+K, Tab, Esc)
//   - ANSI SGR mouse tracking (1006 mode)
//   - Point-and-click A* pathfinding
//   - Context radial menus (right-click)
//   - Fuzzy command palette (Ctrl+K) with success odds
//   - Smart Key (Space/Enter) context resolution
//
// This package MUST NOT import any Layer 1 (transport/crypto/sync/daemon) packages.
package tui

import (
	"errors"
)

// Common errors.
var (
	ErrInvalidKey   = errors.New("invalid key binding")
	ErrModeNotFound = errors.New("mode not found")
	ErrNoAction     = errors.New("no action available")
)

// Mode represents the current TUI mode/focus.
type Mode int

const (
	ModeMapView Mode = iota
	ModeInventory
	ModeSpectrumAnalyzer
	ModeCommandPalette
	ModeRadialMenu
	ModeHelpOverlay
	ModeOnboarding
)

// String returns the mode name.
func (m Mode) String() string {
	switch m {
	case ModeMapView:
		return "Map View"
	case ModeInventory:
		return "Inventory"
	case ModeSpectrumAnalyzer:
		return "RF Spectrum Analyzer"
	case ModeCommandPalette:
		return "Command Palette"
	case ModeRadialMenu:
		return "Radial Menu"
	case ModeHelpOverlay:
		return "Help Overlay"
	case ModeOnboarding:
		return "Onboarding"
	default:
		return "Unknown"
	}
}
