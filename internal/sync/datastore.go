// Package sync provides CRDT state synchronization for Viva Las Mesh Layer 1.
package sync

import (
	"context"

	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
)

// Datastore wraps go-datastore for CRDT persistence.
type Datastore struct {
	ds datastore.Batching
}

// NewDatastore creates a new CRDT datastore wrapper.
func NewDatastore(ds datastore.Batching) *Datastore {
	return &Datastore{ds: ds}
}

// Put stores a CRDT state.
func (d *Datastore) Put(ctx context.Context, key string, crdt CRDT) error {
	state, _, err := crdt.State()
	if err != nil {
		return err
	}
	return d.ds.Put(ctx, datastore.NewKey(key), state)
}

// Get retrieves a CRDT state.
func (d *Datastore) Get(ctx context.Context, key string, crdt CRDT) error {
	data, err := d.ds.Get(ctx, datastore.NewKey(key))
	if err != nil {
		if err == datastore.ErrNotFound {
			return ErrNotFound
		}
		return err
	}
	// Note: Actual deserialization would depend on CRDT type
	_ = data
	return nil
}

// Delete removes a CRDT state.
func (d *Datastore) Delete(ctx context.Context, key string) error {
	return d.ds.Delete(ctx, datastore.NewKey(key))
}

// Query queries CRDT states.
func (d *Datastore) Query(ctx context.Context, q query.Query) (query.Results, error) {
	return d.ds.Query(ctx, q)
}

// Close closes the datastore.
func (d *Datastore) Close() error {
	return d.ds.Close()
}
