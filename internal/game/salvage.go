// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

package game

import (
	"math/rand"
	"time"

	"github.com/opd-ai/vivalasmesh/pkg/types"
)

// SalvageMatrix represents a 6x6 procedural loot generation matrix.
// Rows: material bases (0-5), Columns: wear states (0-5).
type SalvageMatrix [6][6]types.ItemType

// NewSalvageMatrix creates a new salvage matrix with default item types.
// In a full implementation, this would be configured from asset definitions.
func NewSalvageMatrix() SalvageMatrix {
	// Initialize with placeholder item types.
	var m SalvageMatrix
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			// Assign item types based on some logic.
			switch {
			case i == 0 && j == 0:
				m[i][j] = types.ItemConsumable // e.g., Terrible's Cup
			case i == 1 && j == 1:
				m[i][j] = types.ItemKey // e.g., Players Club Card
			case i == 2 && j == 2:
				m[i][j] = types.ItemTool // e.g., SDR Dongle
			case i == 3 && j == 3:
				m[i][j] = types.ItemWeapon // e.g., Knife
			case i == 4 && j == 4:
				m[i][j] = types.ItemCybernetic // e.g., Induction Coil
			case i == 5 && j == 5:
				m[i][j] = types.ItemData // e.g., Elvis Program
			default:
				m[i][j] = types.ItemConsumable // default
			}
		}
	}
	return m
}

// GenerateItem generates an item based on the salvage matrix.
// materialBase: 0-5 representing different material bases (e.g., polymer, metal).
// wearState: 0-5 representing wear state (e.g., pristine, worn, broken).
func (m SalvageMatrix) GenerateItem(materialBase, wearState int) types.ItemType {
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
func (m SalvageMatrix) GenerateRandomItem() types.ItemType {
	// Seed random with current time (in real app, use proper rand source)
	rand.Seed(time.Now().UnixNano())
	materialBase := rand.Intn(6)
	wearState := rand.Intn(6)
	return m.GenerateItem(materialBase, wearState)
}
