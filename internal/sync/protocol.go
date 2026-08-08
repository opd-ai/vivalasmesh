// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
package sync

import (
	"context"
	"sync"
)

// SyncProtocol handles delta synchronization between peers.
type SyncProtocol struct {
	localNodeID string
	datastore   *Datastore
}

// NewSyncProtocol creates a new sync protocol handler.
func NewSyncProtocol(nodeID string, ds *Datastore) *SyncProtocol {
	return &SyncProtocol{
		localNodeID: nodeID,
		datastore:   ds,
	}
}

// ReconciliationHandler handles deterministic state reconciliation for lossy,
// out-of-order packet delivery.
type ReconciliationHandler struct {
	nodeID   string
	protocol *SyncProtocol
	mu       sync.Mutex
	pending  map[string][][]byte // peerID -> list of deltas
}

// NewReconciliationHandler creates a new reconciliation handler.
func NewReconciliationHandler(nodeID string, protocol *SyncProtocol) *ReconciliationHandler {
	return &ReconciliationHandler{
		nodeID:   nodeID,
		protocol: protocol,
		pending:  make(map[string][][]byte),
	}
}

// SendDelta sends a delta to a peer, storing it for potential retransmission.
func (rh *ReconciliationHandler) SendDelta(ctx context.Context, peerID string, crdt CRDT) error {
	// In a real implementation, we would compute the delta and send it over the
	// network, then wait for an acknowledgment. For now, we just store the delta
	// in the pending map to simulate tracking.
	_, version, err := crdt.State()
	if err != nil {
		return err
	}
	delta, err := crdt.Delta(version)
	if err != nil {
		return err
	}
	rh.mu.Lock()
	rh.pending[peerID] = append(rh.pending[peerID], delta)
	rh.mu.Unlock()
	return nil
}

// HandleIncomingDelta processes an incoming delta from a peer, merges it into
// the local CRDT, and returns an acknowledgment delta (to be sent back).
func (rh *ReconciliationHandler) HandleIncomingDelta(ctx context.Context, peerID string, localCRDT, remoteCRDT CRDT) (CRDT, error) {
	// Merge the remote CRDT into the local one.
	changed, err := localCRDT.Merge(remoteCRDT)
	if err != nil {
		return nil, err
	}
	// Persist the merged state if changed.
	if changed {
		_ = rh.protocol.datastore.Put(context.Background(), peerID+"/"+localCRDT.Type(), localCRDT)
	}
	// Return the local CRDT as the acknowledgment (simplified).
	return localCRDT, nil
}

// SyncState synchronizes CRDT state with a peer.
// Returns the merged CRDT or an error.
func (s *SyncProtocol) SyncState(ctx context.Context, peerID string, local, remote CRDT) (CRDT, error) {
	// Get local state
	_, localVersion, err := local.State()
	if err != nil {
		return nil, err
	}

	// Get remote state (in practice, this comes over the network)
	_, remoteVersion, err := remote.State()
	if err != nil {
		return nil, err
	}

	// Compute deltas
	_, err = local.Delta(remoteVersion)
	if err != nil {
		return nil, err
	}

	_, err = remote.Delta(localVersion)
	if err != nil {
		return nil, err
	}

	// Apply remote delta to local
	if _, err := local.Merge(remote); err != nil {
		return nil, err
	}

	// Persist merged state
	if err := s.datastore.Put(ctx, peerID+"/"+local.Type(), local); err != nil {
		return nil, err
	}

	return local, nil
}
