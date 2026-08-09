// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package engine provides the game engine core for Viva Las Mesh Layer 2.
package engine

import (
	"time"
)

// Action represents a staged action in the simulturn system.
type Action struct {
	ID        string
	ActorID   string
	Type      string
	Tier      PriorityTier
	Cost      int    // Action point cost
	Payload   []byte // Serialized action data
	Timestamp time.Time
	ExpiresAt time.Time
	// Attributes holds actor attributes used for tie-breaking in equal-tier collisions.
	Attributes map[string]int // e.g., hustle, hardware, street_smarts
}

// AttributeScore returns a tie-breaker score based on hustle, hardware, and street_smarts.
// Higher score wins.
func (a *Action) AttributeScore() int {
	if a.Attributes == nil {
		return 0
	}
	return a.Attributes["hustle"] + a.Attributes["hardware"] + a.Attributes["street_smarts"]
}

// Validate checks if the action is valid.
func (a *Action) Validate() error {
	if a.Tier < Tier1Divine || a.Tier > Tier5Background {
		return ErrInvalidPriority
	}
	if a.ActorID == "" {
		return ErrInvalidAction
	}
	return nil
}
