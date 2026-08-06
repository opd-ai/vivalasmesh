// Package daemon provides the orchestrator-server daemon for Viva Las Mesh Layer 1.
// This is Layer 1 (P2P Infrastructure) - Background daemon for state replication.
//
// The daemon manages:
//   - Peer discovery & connection management
//   - State replication across transports
//   - CLI for daemon control (start, stop, status, peers)
//
// This package MUST NOT import any Layer 2 (game engine) packages.
package daemon

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	syncpkg "github.com/opd-ai/vivalasmesh/internal/sync"
	"github.com/opd-ai/vivalasmesh/internal/transport"
)

// Common errors.
var (
	ErrDaemonNotRunning = errors.New("daemon not running")
	ErrDaemonRunning    = errors.New("daemon already running")
	ErrPeerNotFound     = errors.New("peer not found")
	ErrTransportClosed  = errors.New("transport closed")
)

// Peer represents a connected peer in the mesh.
type Peer struct {
	ID        string
	Addr      net.Addr
	Transport transport.Transport
	Connected time.Time
	LastSeen  time.Time
	State     *syncpkg.Datastore
}

// DaemonConfig holds daemon configuration.
type DaemonConfig struct {
	NodeID       string
	ListenAddrs  []net.Addr
	Transports   []transport.Transport
	Datastore    *syncpkg.Datastore
	SyncInterval time.Duration
}

// Daemon is the orchestrator-server main loop.
type Daemon struct {
	config    DaemonConfig
	peers     map[string]*Peer
	mu        sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	running   bool
	runningMu sync.Mutex
}

// NewDaemon creates a new orchestrator daemon.
func NewDaemon(config DaemonConfig) *Daemon {
	ctx, cancel := context.WithCancel(context.Background())
	return &Daemon{
		config:  config,
		peers:   make(map[string]*Peer),
		ctx:     ctx,
		cancel:  cancel,
		running: false,
	}
}

// Start starts the daemon main loop.
func (d *Daemon) Start() error {
	d.runningMu.Lock()
	defer d.runningMu.Unlock()

	if d.running {
		return ErrDaemonRunning
	}

	d.running = true
	d.wg.Add(1)
	go d.mainLoop()

	// Start all transports
	for _, t := range d.config.Transports {
		d.wg.Add(1)
		go d.transportLoop(t)
	}

	return nil
}

// Stop stops the daemon gracefully.
func (d *Daemon) Stop() error {
	d.runningMu.Lock()
	defer d.runningMu.Unlock()

	if !d.running {
		return ErrDaemonNotRunning
	}

	d.cancel()
	d.wg.Wait()
	d.running = false
	return nil
}

// Status returns the daemon status.
func (d *Daemon) Status() (map[string]interface{}, error) {
	d.runningMu.Lock()
	defer d.runningMu.Unlock()

	if !d.running {
		return nil, ErrDaemonNotRunning
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	peers := make([]map[string]interface{}, 0, len(d.peers))
	for _, p := range d.peers {
		peers = append(peers, map[string]interface{}{
			"id":        p.ID,
			"addr":      p.Addr.String(),
			"connected": p.Connected,
			"last_seen": p.LastSeen,
			"transport": p.Transport.String(),
		})
	}

	return map[string]interface{}{
		"node_id":       d.config.NodeID,
		"running":       true,
		"peer_count":    len(d.peers),
		"peers":         peers,
		"sync_interval": d.config.SyncInterval,
	}, nil
}

// AddPeer adds a new peer to the daemon.
func (d *Daemon) AddPeer(peer *Peer) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.peers[peer.ID]; exists {
		return errors.New("peer already exists")
	}

	d.peers[peer.ID] = peer
	return nil
}

// RemovePeer removes a peer from the daemon.
func (d *Daemon) RemovePeer(peerID string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.peers[peerID]; !exists {
		return ErrPeerNotFound
	}

	delete(d.peers, peerID)
	return nil
}

// GetPeer returns a peer by ID.
func (d *Daemon) GetPeer(peerID string) (*Peer, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	peer, ok := d.peers[peerID]
	if !ok {
		return nil, ErrPeerNotFound
	}
	return peer, nil
}

// ListPeers returns all connected peers.
func (d *Daemon) ListPeers() []*Peer {
	d.mu.RLock()
	defer d.mu.RUnlock()

	peers := make([]*Peer, 0, len(d.peers))
	for _, p := range d.peers {
		peers = append(peers, p)
	}
	return peers
}

// mainLoop is the daemon's main event loop.
func (d *Daemon) mainLoop() {
	defer d.wg.Done()

	ticker := time.NewTicker(d.config.SyncInterval)
	defer ticker.Stop()

	for {
		select {
		case <-d.ctx.Done():
			return
		case <-ticker.C:
			d.syncPeers()
		}
	}
}

// transportLoop handles incoming frames from a transport.
func (d *Daemon) transportLoop(t transport.Transport) {
	defer d.wg.Done()

	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			frame, addr, err := t.Recv(d.ctx)
			if err != nil {
				// Check if context cancelled
				if d.ctx.Err() != nil {
					return
				}
				// Log error and continue
				continue
			}
			d.handleFrame(t, addr, frame)
		}
	}
}

// handleFrame processes an incoming frame.
func (d *Daemon) handleFrame(t transport.Transport, addr net.Addr, frame []byte) {
	// In a real implementation, this would:
	// 1. Decrypt frame using Noise IK / Double Ratchet
	// 2. Parse frame as CRDT delta
	// 3. Merge into local state
	// 4. Forward to other peers if needed
	_ = frame
	_ = addr
	_ = t
}

// syncPeers synchronizes state with all connected peers.
func (d *Daemon) syncPeers() {
	d.mu.RLock()
	peers := make([]*Peer, 0, len(d.peers))
	for _, p := range d.peers {
		peers = append(peers, p)
	}
	d.mu.RUnlock()

	for _, peer := range peers {
		// In a real implementation, this would:
		// 1. Compute delta since last sync
		// 2. Send delta to peer
		// 2. Receive and merge peer's delta
		_ = peer
	}
}
