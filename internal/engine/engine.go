// Package engine provides the game engine core for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Pulse engine, simulturn, world model.
//
// Implements:
//   - 250ms micro-tick pulse generator
//   - 1000ms macro-pulse coordinator
//   - Action interface with Priority Tier (T1-T5)
//   - Action staging queue (per pulse)
//   - Simulturn resolution (priority sort + tie-breakers)
//   - Attribute system (Hustle, Hardware, Street Smarts, Dignity, etc.)
//
// This package MUST NOT import any Layer 1 (transport/crypto/sync/daemon) packages.
package engine

import (
	"errors"
)

// Common errors.
var (
	ErrInvalidAction   = errors.New("invalid action")
	ErrActionQueueFull = errors.New("action queue full")
	ErrPulseNotRunning = errors.New("pulse engine not running")
	ErrEntityNotFound  = errors.New("entity not found")
	ErrInvalidPriority = errors.New("invalid priority tier")
)

// PriorityTier represents the action priority tier (T1-T5).
// T1: Divine/Spectral King interventions
// T2: Critical system actions (death saves, instant effects)
// T3: Standard actions (move, attack, hack, interact)
// T4: Heavy actions (vault drilling, SCADA hacks)
// T5: Background/environmental (decay, regen, heat)
type PriorityTier int

const (
	Tier1Divine PriorityTier = iota + 1
	Tier2Critical
	Tier3Standard
	Tier4Heavy
	Tier5Background
)

// String returns the tier name.
func (t PriorityTier) String() string {
	switch t {
	case Tier1Divine:
		return "T1:Divine"
	case Tier2Critical:
		return "T2:Critical"
	case Tier3Standard:
		return "T3:Standard"
	case Tier4Heavy:
		return "T4:Heavy"
	case Tier5Background:
		return "T5:Background"
	default:
		return "Unknown"
	}
}
