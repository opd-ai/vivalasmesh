// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"time"
)

// HeatLevel represents the police heat level (0-5 Stars).
type HeatLevel int

const (
	HeatLevel0 HeatLevel = 0 // No heat
	HeatLevel1 HeatLevel = 1 // 1 Star
	HeatLevel2 HeatLevel = 2 // 2 Stars
	HeatLevel3 HeatLevel = 3 // 3 Stars
	HeatLevel4 HeatLevel = 4 // 4 Stars
	HeatLevel5 HeatLevel = 5 // 5 Stars (maximum heat)
)

// HeatStateMachine models the police heat system and dispatch timing.
type HeatStateMachine struct {
	// Current heat level
	Level HeatLevel
	// Base dispatch interval in pulses (how often police check for the player)
	// At heat 0, dispatch every 10 pulses; at heat 5, dispatch every pulse (continuous 250ms micro-ticks)
	baseDispatchInterval int
	// ticksSinceLastDispatch counts pulses since last police dispatch check
	ticksSinceLastDispatch int
	// dispatchTicker is used to simulate pulse ticks (if needed for simulation)
	dispatchTicker *time.Ticker
	// doneChan signals the ticker to stop
	doneChan chan struct{}
}

// NewHeatStateMachine creates a new heat state machine starting at heat 0.
func NewHeatStateMachine() *HeatStateMachine {
	hsm := &HeatStateMachine{
		Level:                  HeatLevel0,
		baseDispatchInterval:   10, // Start with 10 pulses between dispatch checks
		ticksSinceLastDispatch: 0,
		dispatchTicker:         time.NewTicker(250 * time.Millisecond), // Assuming 250ms per pulse
		doneChan:               make(chan struct{}),
	}
	// Start the ticker to simulate pulses
	go hsm.startTicker()
	return hsm
}

// startTicker runs in a goroutine to simulate pulse ticks.
func (hsm *HeatStateMachine) startTicker() {
	for {
		select {
		case <-hsm.dispatchTicker.C:
			hsm.ticksSinceLastDispatch++
			// If we've reached the dispatch interval, trigger a dispatch check and reset counter
			if hsm.ticksSinceLastDispatch >= hsm.baseDispatchInterval {
				// In a real implementation, this would trigger a police dispatch attempt
				hsm.ticksSinceLastDispatch = 0
				// For now, we just log or do nothing; the effect is that dispatch attempts
				// occur more frequently as heat increases (since baseDispatchInterval decreases)
			}
		case <-hsm.doneChan:
			hsm.dispatchTicker.Stop()
			return
		}
	}
}

// IncreaseHeat raises the heat level by the given amount, clamping to 0-5.
// It also updates the base dispatch interval to increase dispatch frequency.
func (hsm *HeatStateMachine) IncreaseHeat(amount int) {
	newLevel := int(hsm.Level) + amount
	if newLevel < 0 {
		newLevel = 0
	}
	if newLevel > 5 {
		newLevel = 5
	}
	hsm.Level = HeatLevel(newLevel)
	// Update dispatch interval: higher heat means more frequent checks
	// Map heat level 0-5 to interval 10-1 (inclusive)
	// Simple linear interpolation: interval = 10 - (heatLevel * 9/5)
	// We'll use integer arithmetic.
	heat := int(hsm.Level)
	hsm.baseDispatchInterval = 10 - (heat*9)/5
	if hsm.baseDispatchInterval < 1 {
		hsm.baseDispatchInterval = 1
	}
}

// DecreaseHeat lowers the heat level by the given amount, clamping to 0-5.
// It also updates the base dispatch interval.
func (hsm *HeatStateMachine) DecreaseHeat(amount int) {
	hsm.IncreaseHeat(-amount) // Reuse IncreaseHeat with negative amount
}

// GetDispatchInterval returns the current number of pulses between dispatch checks.
func (hsm *HeatStateMachine) GetDispatchInterval() int {
	return hsm.baseDispatchInterval
}

// GetHeatLevel returns the current heat level.
func (hsm *HeatStateMachine) GetHeatLevel() HeatLevel {
	return hsm.Level
}

// Stop stops the heat state machine's ticker.
func (hsm *HeatStateMachine) Stop() {
	close(hsm.doneChan)
}
