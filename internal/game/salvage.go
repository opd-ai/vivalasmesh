// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"math/rand"
	"time"

	"github.com/opd-ai/vivalasmesh/pkg/types"
)

// SalvageMatrix represents a 6x6 procedural loot generation matrix.
// Rows: material bases (0-5), Columns: wear states (0-5).
type SalvageMatrix [6][6]*Item

// NewSalvageMatrix creates a new salvage matrix with default items.
// In a full implementation, this would be configured from asset definitions.
func NewSalvageMatrix() SalvageMatrix {
	// Initialize with placeholder items from various eras.
	var m SalvageMatrix
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			// Assign items based on some logic, with era-appropriate examples.
			switch {
			case i == 0 && j == 0:
				// Terrible's Cup - 1953 era item (vintage)
				m[i][j] = NewItem("terribles_cup", "Terrible's Cup", types.ItemConsumable, types.Era1953Atomic,
					"A sun-bleached, cracked, moisture-rotted foam cup fished from a grease-slicked commercial dumpster.", 1)
			case i == 1 && j == 1:
				// Players Club Card - 1993 era item (modern)
				m[i][j] = NewItem("players_club_card", "Players Club Card", types.ItemKey, types.Era1993Corporate,
					"A laminated players club card from a defunct Strip mega-resort.", 5)
			case i == 2 && j == 2:
				// SDR Dongle - 2026 era item (cyber)
				m[i][j] = NewItem("sdr_dongle", "SDR Dongle", types.ItemTool, types.Era2026Cyber,
					"A pocket-sized aluminum receiver tuned to intercept emergency dispatch and hotel maintenance frequencies.", 25)
			case i == 3 && j == 3:
				// Knife - timeless item, but set to 1962 era (classic)
				m[i][j] = NewItem("knife", "Knife", types.ItemWeapon, types.Era1962Syndicate,
					"A standard issue combat knife with pitted rubber grip.", 8)
			case i == 4 && j == 4:
				// Induction Coil - 1993 era item (electronics age)
				m[i][j] = NewItem("induction_coil", "Induction Coil", types.ItemCybernetic, types.Era1993Corporate,
					"A tight spool of enameled wire harvested from an old television monitor chassis.", 15)
			case i == 5 && j == 5:
				// Elvis Program - 2026 era item (cyber/data)
				m[i][j] = NewItem("elvis_program", "Elvis Program", types.ItemData, types.Era2026Cyber,
					"A tiny storage medium recovered from a burned server rack in an abandoned tech incubator.", 50)
			default:
				// Default to a 1953 era consumable
				m[i][j] = NewItem("generic_item", "Generic Item", types.ItemConsumable, types.Era1953Atomic,
					"A generic item for fallback cases.", 0)
			}
		}
	}
	return m
}

// GenerateItem generates an item based on the salvage matrix.
// materialBase: 0-5 representing different material bases (e.g., polymer, metal).
// wearState: 0-5 representing wear state (e.g., pristine, worn, broken).
func (m SalvageMatrix) GenerateItem(materialBase, wearState int) *Item {
	if materialBase < 0 || materialBase > 5 {
		materialBase = 0
	}
	if wearState < 0 || wearState > 5 {
		wearState = 0
	}
	return m[materialBase][wearState]
}

// GenerateRandomItem generates a random item using the salvage matrix.
// This simulates procuring a random salvage item.
func (m SalvageMatrix) GenerateRandomItem() *Item {
	// Seed random with current time (in real app, use proper rand source)
	rand.Seed(time.Now().UnixNano())
	materialBase := rand.Intn(6)
	wearState := rand.Intn(6)
	return m.GenerateItem(materialBase, wearState)
}
