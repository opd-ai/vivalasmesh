// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
package sync

import (
	"github.com/cbergoon/merkletree"
)

// MerkleTree provides Merkle tree verification for state.
type MerkleTree struct {
	tree *merkletree.MerkleTree
}

// NewMerkleTree creates a new Merkle tree from content.
func NewMerkleTree(contents []merkletree.Content) (*MerkleTree, error) {
	tree, err := merkletree.NewTree(contents)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{tree: tree}, nil
}

// Root returns the Merkle root hash.
func (m *MerkleTree) Root() []byte {
	return m.tree.MerkleRoot()
}

// Verify verifies a proof for a given content.
func (m *MerkleTree) Verify(content merkletree.Content) (bool, error) {
	return m.tree.VerifyContent(content)
}
