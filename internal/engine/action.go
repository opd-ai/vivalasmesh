// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package engine provides the game engine core for Viva Las Mesh Layer 2.
package engine

import (
	"time"
)

// Action represents a staged action in the simulturn system.
type Action struct {
	ID         string
	ActorID    string
	Type       string
	Tier       PriorityTier
	Cost       int    // Action point cost
	Payload    []byte // Serialized action data
	Timestamp  time.Time
	ExpiresAt  time.Time
	Attributes map[string]int // Required attribute checks
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
