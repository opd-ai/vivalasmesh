// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"time"
)

// EscapologyPhase represents the current phase of the Desert Pit Escapology.
type EscapologyPhase int

const (
	PhaseTrunkTransport EscapologyPhase = iota // Phase 1: Trunk Transport (12s)
	PhaseGraveDigging                          // Phase 2: Grave Digging (15s)
	PhaseExecutionDraw                         // Phase 3: Execution Draw (3s)
	PhaseDesertStandoff                        // Phase 4: Desert Standoff (5s)
	PhaseCompleted                             // Escapology completed successfully
	PhaseFailed                                // Escapology failed (timeout, wrong input, etc.)
)

// EscapologyStateMachine manages the state of the Desert Pit Escapology.
// This is a real-time sequence where the player must perform inputs within time limits.
type EscapologyStateMachine struct {
	// Current phase
	Phase EscapologyPhase
	// Start time of the current phase
	StartTime time.Time
	// Time limit for the current phase in seconds
	TimeLimit int
	// Indicates if the escapology is active
	Active bool
	// Failed flag
	Failed bool
	// Completed flag
	Completed bool
	// Input channel for receiving player inputs (simplified)
	inputChan chan struct{}
	// Ticker for checking timeouts
	ticker *time.Ticker
	// Done channel to stop the ticker
	doneChan chan struct{}
}

// NewEscapologyStateMachine creates a new escapology state machine.
func NewEscapologyStateMachine() *EscapologyStateMachine {
	esm := &EscapologyStateMachine{
		Phase:     PhaseTrunkTransport,
		TimeLimit: 12, // Default for first phase
		Active:    false,
		Failed:    false,
		Completed: false,
		inputChan: make(chan struct{}, 10),         // Buffered channel for inputs
		ticker:    time.NewTicker(1 * time.Second), // Check every second
		doneChan:  make(chan struct{}),
	}
	// Start the ticker to monitor time
	go esm.startTicker()
	return esm
}

// Start begins the escapology, setting the phase to Trunk Transport.
func (esm *EscapologyStateMachine) Start() {
	esm.Phase = PhaseTrunkTransport
	esm.StartTime = time.Now()
	esm.TimeLimit = 12
	esm.Active = true
	esm.Failed = false
	esm.Completed = false
	// Reset input channel
	esm.inputChan = make(chan struct{}, 10)
}

// startTicker runs in a goroutine to check for timeouts.
func (esm *EscapologyStateMachine) startTicker() {
	for {
		select {
		case <-esm.ticker.C:
			elapsed := int(time.Since(esm.StartTime).Seconds())
			if elapsed >= esm.TimeLimit {
				// Timeout for current phase
				esm.Fail()
				return
			}
		case <-esm.doneChan:
			esm.ticker.Stop()
			return
		}
	}
}

// InputReceived signals that the player has provided the correct input for the current phase.
// This should be called when the player performs the required action (e.g., presses a key at the right time).
func (esm *EscapologyStateMachine) InputReceived() {
	if !esm.Active {
		return
	}
	// Signal that input was received
	select {
	case esm.inputChan <- struct{}{}:
	default:
		// Drop if channel is full (shouldn't happen with buffering)
	}
	// Check if the input was timely (we rely on the ticker to have not timed out yet)
	// If we reach here before timeout, we can advance to next phase
	esm.AdvancePhase()
}

// AdvancePhase moves to the next phase if the current phase was completed successfully.
func (esm *EscapologyStateMachine) AdvancePhase() {
	if !esm.Active {
		return
	}
	switch esm.Phase {
	case PhaseTrunkTransport:
		esm.Phase = PhaseGraveDigging
		esm.StartTime = time.Now()
		esm.TimeLimit = 15
	case PhaseGraveDigging:
		esm.Phase = PhaseExecutionDraw
		esm.StartTime = time.Now()
		esm.TimeLimit = 3
	case PhaseExecutionDraw:
		esm.Phase = PhaseDesertStandoff
		esm.StartTime = time.Now()
		esm.TimeLimit = 5
	case PhaseDesertStandoff:
		esm.Phase = PhaseCompleted
		esm.Active = false
		esm.Completed = true
	case PhaseCompleted, PhaseFailed:
		// Already finished
	}
}

// Fail marks the escapology as failed (e.g., timeout, wrong input).
func (esm *EscapologyStateMachine) Fail() {
	esm.Active = false
	esm.Failed = true
	// Stop the ticker
	esm.doneChan <- struct{}{}
}

// IsActive returns true if the escapology is currently in progress.
func (esm *EscapologyStateMachine) IsActive() bool {
	return esm.Active
}

// IsCompleted returns true if the escapology was completed successfully.
func (esm *EscapologyStateMachine) IsCompleted() bool {
	return esm.Completed
}

// IsFailed returns true if the escapology failed.
func (esm *EscapologyStateMachine) IsFailed() bool {
	return esm.Failed
}

// GetCurrentPhase returns the current phase of the escapology.
func (esm *EscapologyStateMachine) GetCurrentPhase() EscapologyPhase {
	return esm.Phase
}
