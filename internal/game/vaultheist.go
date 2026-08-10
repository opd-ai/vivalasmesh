// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"time"
)

// VaultHeistStage represents the current stage of a vault heist.
type VaultHeistStage int

const (
	StageSurveillanceBypass  VaultHeistStage = iota // Stage 1: Surveillance Bypass
	StageGuardNeutralization                        // Stage 2: Guard Neutralization
	StageVaultUnlock                                // Stage 3: Vault Unlock
	StageCashExtraction                             // Stage 4: Cash Extraction
	StageCompleted                                  // Heist completed
	StageFailed                                     // Heist failed (caught, alarm triggered, etc.)
)

// VaultHeistStateMachine manages the state of a vault heist.
type VaultHeistStateMachine struct {
	// Current stage of the heist
	Stage VaultHeistStage
	// Timestamp when the heist started (or last stage transition)
	StartTime time.Time
	// Indicates if the heist is currently active
	Active bool
	// Failed flag if the heist failed
	Failed bool
	// Completed flag if the heist succeeded
	Completed bool
}

// NewVaultHeistStateMachine creates a new vault heist state machine, ready to start.
func NewVaultHeistStateMachine() *VaultHeistStateMachine {
	return &VaultHeistStateMachine{
		Stage:     StageSurveillanceBypass,
		Active:    false,
		Failed:    false,
		Completed: false,
	}
}

// Start begins the vault heist, setting the stage to Surveillance Bypass.
func (vhm *VaultHeistStateMachine) Start() {
	vhm.Stage = StageSurveillanceBypass
	vhm.StartTime = time.Now()
	vhm.Active = true
	vhm.Failed = false
	vhm.Completed = false
}

// AttemptSurveillanceBypass attempts to bypass surveillance (cameras, motion sensors, etc.).
// Returns true if successful, allowing transition to the next stage.
// In a real implementation, this would involve skill checks, equipment, etc.
func (vhm *VaultHeistStateMachine) AttemptSurveillanceBypass() bool {
	if !vhm.Active || vhm.Stage != StageSurveillanceBypass {
		return false
	}
	// Placeholder: assume success if we have adequate skills (to be implemented)
	// For now, we'll return true to allow progression (can be made more complex later)
	return true
}

// AttemptGuardNeutralization attempts to neutralize guards without raising alarm.
// Returns true if successful.
func (vhm *VaultHeistStateMachine) AttemptGuardNeutralization() bool {
	if !vhm.Active || vhm.Stage != StageGuardNeutralization {
		return false
	}
	// Placeholder
	return true
}

// AttemptVaultUnlock attempts to unlock the vault using hacking, gadgets, or force.
// Returns true if successful.
func (vhm *VaultHeistStateMachine) AttemptVaultUnlock() bool {
	if !vhm.Active || vhm.Stage != StageVaultUnlock {
		return false
	}
	// Placeholder
	return true
}

// AttemptCashExtraction attempts to extract cash from the vault and exit.
// Returns true if successful, completing the heist.
func (vhm *VaultHeistStateMachine) AttemptCashExtraction() bool {
	if !vhm.Active || vhm.Stage != StageCashExtraction {
		return false
	}
	// Placeholder
	return true
}

// AdvanceStage moves the heist to the next stage if the current stage action was successful.
// This should be called after a successful action attempt.
func (vhm *VaultHeistStateMachine) AdvanceStage() {
	if !vhm.Active {
		return
	}
	switch vhm.Stage {
	case StageSurveillanceBypass:
		vhm.Stage = StageGuardNeutralization
	case StageGuardNeutralization:
		vhm.Stage = StageVaultUnlock
	case StageVaultUnlock:
		vhm.Stage = StageCashExtraction
	case StageCashExtraction:
		vhm.Stage = StageCompleted
		vhm.Active = false
		vhm.Completed = true
	case StageCompleted, StageFailed:
		// Already finished
	}
}

// FailHeist marks the heist as failed (e.g., alarm triggered, caught by guards).
func (vhm *VaultHeistStateMachine) FailHeist() {
	vhm.Active = false
	vhm.Failed = true
	vhm.Stage = StageFailed
}

// IsActive returns true if the heist is currently in progress.
func (vhm *VaultHeistStateMachine) IsActive() bool {
	return vhm.Active
}

// IsCompleted returns true if the heist was completed successfully.
func (vhm *VaultHeistStateMachine) IsCompleted() bool {
	return vhm.Completed
}

// IsFailed returns true if the heist failed.
func (vhm *VaultHeistStateMachine) IsFailed() bool {
	return vhm.Failed
}

// GetCurrentStage returns the current stage of the heist.
func (vhm *VaultHeistStateMachine) GetCurrentStage() VaultHeistStage {
	return vhm.Stage
}
