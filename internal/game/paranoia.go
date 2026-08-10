// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// ParanoiaIndex represents the player's paranoia level, calculated from various factors.
// Higher values indicate greater paranoia, leading to visual and gameplay effects.
type ParanoiaIndex int

// CalculateParanoiaIndex computes the paranoia index based on THC dosage, alcohol level,
// metabolic ticks, heat level, and street smarts.
// The formula is a simplified linear combination with weights.
// In a real implementation, this could be more complex, with non-linear effects and thresholds.
func CalculateParanoiaIndex(thcDosage, alcoholLevel, metabolicTicks int, heatLevel HeatLevel, streetSmarts int) ParanoiaIndex {
	// Weights for each factor (empirically determined for game balance)
	const (
		weightTHC          = 2 // Each THC dosage unit adds 2 paranoia points
		weightAlcohol      = 1 // Each alcohol unit adds 1 paranoia point
		weightMetabolic    = 1 // Each metabolic tick adds 1 paranoia point
		weightHeat         = 5 // Each heat level adds 5 paranoia points (heat significantly increases paranoia)
		weightStreetSmarts = 3 // Each street smarts point reduces paranoia by 3 (street smarts helps cope)
	)

	// Calculate raw paranoia score
	raw := thcDosage*weightTHC +
		alcoholLevel*weightAlcohol +
		metabolicTicks*weightMetabolic +
		int(heatLevel)*weightHeat -
		streetSmarts*weightStreetSmarts

	// Clamp to a reasonable range, e.g., 0 to 100
	if raw < 0 {
		raw = 0
	}
	if raw > 100 {
		raw = 100
	}
	return ParanoiaIndex(raw)
}

// GetParanoiaState returns a description of the paranoia state based on the index.
// This can be used to trigger specific visual or gameplay effects.
func GetParanoiaState(index ParanoiaIndex) string {
	switch {
	case index >= 1 && index <= 5:
		return "Smooth Synergy"
	case index >= 6 && index <= 12:
		return "Greening Out"
	case index >= 13 && index <= 20:
		return "Bat Country"
	case index >= 21:
		return "Existential Dread"
	default:
		return "Normal"
	}
}
