// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// MorgueWakeState represents the state of a player who has awakened in the morgue.
// This is a severe penalty for failing to escape dangerous situations.
type MorgueWakeState struct {
	// FLAG_TOE_TAG_ACTIVE indicates the player is marked as deceased (toe tag active)
	FLAG_TOE_TAG_ACTIVE bool
	// PhysicalInventory represents the player's physical items (weapons, tools, etc.)
	// In the morgue wake state, all physical inventory is stripped
	PhysicalInventory []string
	// PoliceAggression represents the current police aggression level (0-5 stars)
	// In the morgue wake state, police aggression is reset due to being "dead"
	PoliceAggression int
	// Other effects: disorientation, stat penalties, etc.
	Health       int
	Buzz         int
	Hydration    int
	StreetSmarts int
	Dignity      int
	Hardware     int
}

// HandleMorgueWake processes the morgue wake state for a player.
// It sets FLAG_TOE_TAG_ACTIVE, strips physical inventory, and resets police aggression.
// It also applies stat penalties due to the traumatic experience.
func HandleMorgueWake(player *PlayerState) MorgueWakeState {
	// In a real implementation, this would modify the actual player state.
	// For now, we return a new MorgueWakeState representing the outcome.
	return MorgueWakeState{
		FLAG_TOE_TAG_ACTIVE: true,
		PhysicalInventory:   []string{}, // All inventory stripped
		PoliceAggression:    0,          // Police aggression reset (player presumed dead)
		// Apply stat penalties: reduced health, buzz, etc.
		Health:       max(0, player.Health-20),      // Significant health reduction
		Buzz:         max(0, player.Buzz-5),         // Reduced buzz tolerance
		Hydration:    max(0, player.Hydration-30),   // Dehydration
		StreetSmarts: max(0, player.StreetSmarts-5), // Disorientation reduces street smarts
		Dignity:      max(0, player.Dignity-10),     // Loss of dignity
		Hardware:     max(0, player.Hardware-5),     // Fine motor skills impaired
	}
}

// PlayerState represents the player's current state (simplified for this example).
type PlayerState struct {
	Health       int
	Buzz         int
	Hydration    int
	StreetSmarts int
	Dignity      int
	Hardware     int
	// In a real implementation, this would include inventory, flags, etc.
	Inventory []string
	Flags     map[string]bool
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
