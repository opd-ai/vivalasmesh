// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// Faction represents the major factions in the game.
type Faction string

const (
	FactionAEC           Faction = "AEC"            // Atomic Energy Commission
	FactionChicagoOutfit Faction = "CHICAGO_OUTFIT" // Chicago Outfit
	FactionCorporate     Faction = "CORPORATE"      // Corporate Moguls
	FactionDarknet       Faction = "DARKNET"        // Darknet Hackers
	FactionMetroPD       Faction = "METRO_PD"       // Metro Police Department
)

// FactionDisposition represents a faction's disposition towards the player,
// ranging from -100 (hostile) to +100 (friendly).
type FactionDisposition int

// GetInitialDisposition returns the starting disposition for each faction.
// In a real implementation, this might be influenced by starting era, background, etc.
func GetInitialDisposition() map[Faction]FactionDisposition {
	return map[Faction]FactionDisposition{
		FactionAEC:           0, // Neutral start
		FactionChicagoOutfit: 0, // Neutral start
		FactionCorporate:     0, // Neutral start
		FactionDarknet:       0, // Neutral start
		FactionMetroPD:       0, // Neutral start
	}
}

// UpdateDisposition updates a faction's disposition based on an action.
// This is a simplified placeholder; in reality, each action would affect
// different factions based on the action type and context.
func UpdateDisposition(current map[Faction]FactionDisposition, faction Faction, change int) map[Faction]FactionDisposition {
	if current == nil {
		current = GetInitialDisposition()
	}
	if val, exists := current[faction]; exists {
		newVal := val + FactionDisposition(change)
		// Clamp between -100 and +100
		if newVal < -100 {
			newVal = -100
		}
		if newVal > 100 {
			newVal = 100
		}
		current[faction] = newVal
	}
	return current
}
