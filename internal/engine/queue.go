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

// Resolve sorts actions by priority tier and resolves conflicts.
// Returns actions in resolution order.
func (q *ActionQueue) Resolve() []*Action {
	// Sort by tier (ascending = higher priority first), then by timestamp
	// In a real implementation, this would be more complex with tie-breakers
	result := make([]*Action, len(q.actions))
	copy(result, q.actions)

	// Simple bubble sort for now (small N)
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j].Tier > result[j+1].Tier ||
				(result[j].Tier == result[j+1].Tier && result[j].Timestamp.After(result[j+1].Timestamp)) {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	// Clear queue after resolution
	q.actions = q.actions[:0]
	return result
}

// Size returns the number of staged actions.
func (q *ActionQueue) Size() int {
	return len(q.actions)
}
