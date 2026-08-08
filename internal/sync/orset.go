// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
package sync

import (
	"encoding/json"
	"time"
)

// ORSet is an Observed-Remove Set CRDT.
// Suitable for collections where add/remove operations must converge.
type ORSet struct {
	elements map[string]*ORSetElement
	nodeID   string
}

type ORSetElement struct {
	Value     []byte
	AddTag    string
	RemoveTag string
	Timestamp time.Time
}

// NewORSet creates a new OR-Set.
func NewORSet(nodeID string) *ORSet {
	return &ORSet{
		elements: make(map[string]*ORSetElement),
		nodeID:   nodeID,
	}
}

// Add adds an element to the set.
func (s *ORSet) Add(value []byte) {
	tag := generateTag(s.nodeID)
	key := string(value)
	s.elements[key] = &ORSetElement{
		Value:     value,
		AddTag:    tag,
		Timestamp: time.Now(),
	}
}

// Remove removes an element from the set.
func (s *ORSet) Remove(value []byte) {
	key := string(value)
	if elem, ok := s.elements[key]; ok {
		elem.RemoveTag = generateTag(s.nodeID)
	}
}

// Contains checks if an element is in the set.
func (s *ORSet) Contains(value []byte) bool {
	key := string(value)
	elem, ok := s.elements[key]
	if !ok {
		return false
	}
	return elem.RemoveTag == ""
}

// Elements returns all current elements.
func (s *ORSet) Elements() [][]byte {
	var result [][]byte
	for _, elem := range s.elements {
		if elem.RemoveTag == "" {
			result = append(result, elem.Value)
		}
	}
	return result
}

// Merge implements CRDT interface.
func (s *ORSet) Merge(other CRDT) (bool, error) {
	o, ok := other.(*ORSet)
	if !ok {
		return false, ErrInvalidDelta
	}

	changed := false
	for key, elem := range o.elements {
		if local, ok := s.elements[key]; !ok {
			s.elements[key] = elem
			changed = true
		} else if elem.Timestamp.After(local.Timestamp) {
			s.elements[key] = elem
			changed = true
		}
	}
	return changed, nil
}

// Delta implements CRDT interface.
func (s *ORSet) Delta(version interface{}) ([]byte, error) {
	v, ok := version.(time.Time)
	if !ok {
		return nil, ErrInvalidDelta
	}

	var deltas []*ORSetElement
	for _, elem := range s.elements {
		if elem.Timestamp.After(v) {
			deltas = append(deltas, elem)
		}
	}

	if len(deltas) == 0 {
		return nil, nil
	}
	return json.Marshal(deltas)
}

// State implements CRDT interface.
func (s *ORSet) State() ([]byte, interface{}, error) {
	var elements []*ORSetElement
	for _, elem := range s.elements {
		elements = append(elements, elem)
	}
	data, _ := json.Marshal(elements)
	return data, time.Now(), nil
}

// Type implements CRDT interface.
func (s *ORSet) Type() string {
	return "ORSet"
}

func generateTag(nodeID string) string {
	return nodeID + "-" + time.Now().Format("20060102150405.000000000")
}
