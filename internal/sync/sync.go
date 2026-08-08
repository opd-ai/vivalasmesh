// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
// This is Layer 1 (P2P Infrastructure) - CRDT, Merkle tree, state convergence.
//
// Implements:
//   - CRDT interface (Merge, Delta, State)
//   - LWW-Register for simple values
//   - OR-Set for collections
//   - Merkle tree for state verification (cbergoon/merkletree)
//   - go-datastore persistence layer
//   - Delta synchronization protocol
//
// This package MUST NOT import any Layer 2 (game engine) packages.
package sync

import (
	"encoding/json"
	"errors"
	"time"
)

// Common errors.
var (
	ErrNotFound        = errors.New("key not found")
	ErrMergeConflict   = errors.New("merge conflict: concurrent modification")
	ErrInvalidDelta    = errors.New("invalid delta format")
	ErrDatastoreClosed = errors.New("datastore is closed")
)

// CRDT defines the interface for Conflict-free Replicated Data Types.
// All implementations must be thread-safe and convergent.
type CRDT interface {
	// Merge merges another CRDT's state into this one.
	// Returns true if state changed, false if already converged.
	Merge(other CRDT) (bool, error)

	// Delta returns the delta since the given version.
	// Version is an opaque token returned by State().
	Delta(version interface{}) ([]byte, error)

	// State returns the current state and a version token.
	State() ([]byte, interface{}, error)

	// Type returns the CRDT type identifier.
	Type() string
}

// LWWRegister is a Last-Writer-Wins Register CRDT.
// Suitable for single values where last write wins (timestamp-based).
type LWWRegister struct {
	value     []byte
	timestamp time.Time
	nodeID    string
}

// NewLWWRegister creates a new LWW register.
func NewLWWRegister(nodeID string) *LWWRegister {
	return &LWWRegister{
		nodeID: nodeID,
	}
}

// Set sets the register value with current timestamp.
func (r *LWWRegister) Set(value []byte) {
	r.value = value
	r.timestamp = time.Now()
}

// Get returns the current value.
func (r *LWWRegister) Get() []byte {
	return r.value
}

// Merge implements CRDT interface.
func (r *LWWRegister) Merge(other CRDT) (bool, error) {
	o, ok := other.(*LWWRegister)
	if !ok {
		return false, ErrInvalidDelta
	}

	if o.timestamp.After(r.timestamp) {
		r.value = o.value
		r.timestamp = o.timestamp
		return true, nil
	}
	return false, nil
}

// Delta implements CRDT interface.
func (r *LWWRegister) Delta(version interface{}) ([]byte, error) {
	v, ok := version.(time.Time)
	if !ok {
		return nil, ErrInvalidDelta
	}
	if r.timestamp.After(v) {
		data, _ := json.Marshal(struct {
			Value     []byte
			Timestamp time.Time
			NodeID    string
		}{
			Value:     r.value,
			Timestamp: r.timestamp,
			NodeID:    r.nodeID,
		})
		return data, nil
	}
	return nil, nil
}

// State implements CRDT interface.
func (r *LWWRegister) State() ([]byte, interface{}, error) {
	data, _ := json.Marshal(struct {
		Value     []byte
		Timestamp time.Time
		NodeID    string
	}{
		Value:     r.value,
		Timestamp: r.timestamp,
		NodeID:    r.nodeID,
	})
	return data, r.timestamp, nil
}

// Type implements CRDT interface.
func (r *LWWRegister) Type() string {
	return "LWWRegister"
}
