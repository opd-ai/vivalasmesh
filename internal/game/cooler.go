// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"time"
)

// CosmicCooler represents "The Cosmic Cooler" roaming NPC that drains player luck.
type CosmicCooler struct {
	// isActive indicates whether the Cooler is currently active in the game world
	isActive bool
	// LastUpdateTime tracks when the Cooler last drained luck
	LastUpdateTime time.Time
	// DrainInterval is how often the Cooler attempts to drain luck
	DrainInterval time.Duration
	// EffectRadius is the radius within which players are affected (in game units)
	EffectRadius int
}

// NewCosmicCooler creates a new Cosmic Cooler NPC.
func NewCosmicCooler() *CosmicCooler {
	return &CosmicCooler{
		isActive:       false,
		LastUpdateTime: time.Now(),
		DrainInterval:  30 * time.Second, // Check every 30 seconds
		EffectRadius:   50,               // Affects players within 50 units
	}
}

// Start activates the Cosmic Cooler in the game world.
func (cc *CosmicCooler) Start() {
	cc.isActive = true
	cc.LastUpdateTime = time.Now()
}

// Stop deactivates the Cosmic Cooler.
func (cc *CosmicCooler) Stop() {
	cc.isActive = false
}

// ApplyLuckDrain drains 50% of the player's current luck (if within effect radius).
// This is a simplified simulation; in a real implementation, this would involve
// CRDT state updates to propagate the luck drain across the network.
func (cc *CosmicCooler) ApplyLuckDrain(currentLuck int) int {
	if !cc.isActive {
		return currentLuck
	}
	// Check if enough time has passed since last drain
	if time.Since(cc.LastUpdateTime) < cc.DrainInterval {
		return currentLuck // Not yet time to drain again
	}
	// Drain 50% of luck
	drained := currentLuck / 2
	cc.LastUpdateTime = time.Now()
	return drained // Return the remaining luck after drain
}

// IsActive returns true if the Cooler is currently active.
func (cc *CosmicCooler) IsActive() bool {
	return cc.isActive
}

// SetEffectRadius sets the radius within which players are affected by the Cooler.
func (cc *CosmicCooler) SetEffectRadius(radius int) {
	cc.EffectRadius = radius
}

// GetEffectRadius returns the effect radius.
func (cc *CosmicCooler) GetEffectRadius() int {
	return cc.EffectRadius
}
