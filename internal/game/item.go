// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"github.com/opd-ai/vivalasmesh/pkg/types"
)

// Item represents a game item with properties beyond just its type.
type Item struct {
	ID          string
	Name        string
	ItemType    types.ItemType
	Era         types.Era // The era this item is from (its "tech level")
	Description string
	Value       int // Value in credits
}

// NewItem creates a new item with the specified properties.
func NewItem(id, name string, itemType types.ItemType, era types.Era, description string, value int) *Item {
	return &Item{
		ID:          id,
		Name:        name,
		ItemType:    itemType,
		Era:         era,
		Description: description,
		Value:       value,
	}
}
