// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

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
	"sync"
	"time"
)

// Engine represents the core game engine.
type Engine struct {
	pulseEngine *PulseEngine
	running     bool
	mu          sync.Mutex
}

// NewEngine creates a new game engine with the given micro and macro tick durations.
func NewEngine() *Engine {
	return &Engine{
		pulseEngine: NewPulseEngine(
			250*time.Millisecond,
			1000*time.Millisecond,
			10*time.Second,
			60*time.Second,
			1024),
	}
}

// Start starts the engine and its pulse loop.
func (e *Engine) Start() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.running {
		return ErrPulseNotRunning // Reuse error, or define a new one
	}
	if err := e.pulseEngine.Start(); err != nil {
		return err
	}
	// Set up pulse handlers
	e.pulseEngine.SetMicroTickHandler(func(tick int) {
		// Handle 250ms micro-tick: real-time movement interpolation, input polling, sub-pixel redraws
		// TODO: Implement actual game logic for micro-tick
	})
	e.pulseEngine.SetMacroPulseHandler(func() {
		// Handle 1000ms macro-pulse: room state updates, guard pathfinding execution, SCADA timers, packet dispatching
		// TODO: Implement actual game logic for macro-pulse
	})
	e.pulseEngine.SetMetaPulseHandler(func() {
		// Handle 10.0s meta-pulse: metabolic decay, hydration consumption, police heat decay, mob interest accrual
		// TODO: Implement actual game logic for meta-pulse
	})
	e.pulseEngine.SetEonPulseHandler(func() {
		// Handle 60.0s eon-pulse: spectral apparition checks, cross-network casino jackpot pool drift
		// TODO: Implement actual game logic for eon-pulse
	})
	e.running = true
	return nil
}

// Stop stops the engine and its pulse loop.
func (e *Engine) Stop() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if !e.running {
		return ErrPulseNotRunning
	}
	if err := e.pulseEngine.Stop(); err != nil {
		return err
	}
	e.running = false
	return nil
}

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
