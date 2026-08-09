// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package engine provides the game engine core for Viva Las Mesh Layer 2.
package engine

// ActionQueue holds staged actions for a pulse.
type ActionQueue struct {
	actions []*Action
	maxSize int
}

// NewActionQueue creates a new action queue.
func NewActionQueue(maxSize int) *ActionQueue {
	return &ActionQueue{
		actions: make([]*Action, 0, maxSize),
		maxSize: maxSize,
	}
}

// Stage adds an action to the queue.
func (q *ActionQueue) Stage(action *Action) error {
	if len(q.actions) >= q.maxSize {
		return ErrActionQueueFull
	}
	if err := action.Validate(); err != nil {
		return err
	}
	q.actions = append(q.actions, action)
	return nil
}

// Resolve sorts actions by priority tier and resolves conflicts using simultaneous resolution rules.
// Returns actions in resolution order.
func (q *ActionQueue) Resolve() []*Action {
	// Sort by tier (ascending = higher priority first), then by tie-breaker score, then by timestamp
	result := make([]*Action, len(q.actions))
	copy(result, q.actions)

	// Simple bubble sort for now (small N)
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			a := result[j]
			b := result[j+1]
			if a.Tier > b.Tier {
				result[j], result[j+1] = b, a
			} else if a.Tier == b.Tier {
				// Tie-breaker: compare attribute score (hustle+hardware+street_smarts), higher wins
				scoreA := a.AttributeScore()
				scoreB := b.AttributeScore()
				if scoreA < scoreB {
					result[j], result[j+1] = b, a
				} else if scoreA == scoreB {
					// If still tied, use timestamp (older first)
					if a.Timestamp.After(b.Timestamp) {
						result[j], result[j+1] = b, a
					}
				}
			}
		}
	}

	// Simultaneous resolution solver:
	// Process actions in priority order and resolve conflicts.
	// For simplicity, we treat actions with the same Type and Payload as conflicting.
	// In case of conflict, only the highest priority action (first in sorted order) is kept.
	// This resolver can be extended to handle specific action types (e.g., move, attack, pickup)
	// and apply appropriate simultaneous resolution rules (e.g., both attacks succeed, move collisions cause bounceback).
	var resolved []*Action
	seen := make(map[string]struct{})
	for _, a := range result {
		key := string(a.Type) + ":" + string(a.Payload)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			resolved = append(resolved, a)
			// Log conflict detection for debugging (can be removed in production)
			if len(resolved) > 1 {
				// Check if this action conflicts with any previously resolved action
				// For now, we just note that we are processing another action of same Type/Payload
				// but we skip duplicates due to the seen map.
				// In a more advanced solver, we would apply conflict resolution logic here.
			}
		} else {
			// Duplicate action detected (same Type and Payload). This indicates a potential
			// concurrent action conflict (e.g., two pickups of the same item, two attacks on the same target).
			// According to simultaneous resolution rules, we skip the duplicate action.
			// TODO: Implement specific conflict resolution based on action type.
			// For now, we log and skip.
			// log.Printf("Conflict detected: duplicate action Type=%s, Payload=%v", a.Type, a.Payload)
		}
	}

	// Clear queue after resolution
	q.actions = q.actions[:0]
	return resolved
}

// Size returns the number of staged actions.
func (q *ActionQueue) Size() int {
	return len(q.actions)
}
