// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

package engine

import (
	"reflect"
	"testing"
	"time"
)

func TestResolve_SimultaneousResolution(t *testing.T) {
	// Create a queue with capacity 10
	q := NewActionQueue(10)

	// Create actions with same Type and Payload (should conflict)
	act1 := &Action{
		ID:         "1",
		ActorID:    "actor1",
		Type:       "move",
		Tier:       Tier3Standard,
		Payload:    []byte(`{"target":"tile10,5"}`),
		Timestamp:  time.Now(),
		Attributes: map[string]int{"hustle": 5, "hardware": 5, "street_smarts": 5},
	}
	act2 := &Action{
		ID:         "2",
		ActorID:    "actor2",
		Type:       "move",
		Tier:       Tier3Standard,                                                  // same tier
		Payload:    []byte(`{"target":"tile10,5"}`),                                // same payload -> conflict
		Timestamp:  time.Now().Add(-time.Second),                                   // older
		Attributes: map[string]int{"hustle": 3, "hardware": 3, "street_smarts": 3}, // lower attribute score
	}
	act3 := &Action{
		ID:         "3",
		ActorID:    "actor3",
		Type:       "attack",
		Tier:       Tier3Standard,
		Payload:    []byte(`{"target":"actor2"}`),
		Timestamp:  time.Now(),
		Attributes: map[string]int{"hustle": 10, "hardware": 10, "street_smarts": 10}, // higher attribute score
	}

	// Stage actions
	if err := q.Stage(act1); err != nil {
		t.Fatalf("Failed to stage act1: %v", err)
	}
	if err := q.Stage(act2); err != nil {
		t.Fatalf("Failed to stage act2: %v", err)
	}
	if err := q.Stage(act3); err != nil {
		t.Fatalf("Failed to stage act3: %v", err)
	}

	// Resolve
	resolved := q.Resolve()

	// Expect only two actions resolved: act3 (higher attribute score) and act1 (older timestamp? Actually act1 and act2 same tier, act1 has higher attribute score? Wait act1 attributes sum 15, act2 sum 9, so act1 should win over act2 due to attribute score.
	// However, act3 has same tier as act1 and act2, but higher attribute score (30) so act3 should be first.
	// Then between act1 and act2, act1 has higher attribute score, so act1 should be second.
	// Thus resolved order should be [act3, act1] (act2 omitted due to conflict with act1? Actually act2 conflicts with act1 (same Type and Payload), so only one of them should be kept.
	// Since we process in sorted order, the sorted order should be: act3 (highest attr), act1 (next attr), act2 (lowest attr).
	// Then we keep first of each Type/Payload group.
	// Group "move":{"target":"tile10,5"} -> act1 and act2, keep act1 (higher attr).
	// Group "attack":{"target":"actor2"} -> act3, keep act3.
	// So resolved should be [act3, act1].

	if len(resolved) != 2 {
		t.Fatalf("Expected 2 resolved actions, got %d", len(resolved))
	}
	if !reflect.DeepEqual(resolved[0], act3) {
		t.Fatalf("Expected first resolved action to be act3, got %v", resolved[0])
	}
	if !reflect.DeepEqual(resolved[1], act1) {
		t.Fatalf("Expected second resolved action to be act1, got %v", resolved[1])
	}
}
