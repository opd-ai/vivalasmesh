// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
package sync

import (
	"context"
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
