// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package transport provides the transport abstraction layer for Viva Las Mesh.
// This is Layer 1 (P2P Infrastructure) - Real-world multi-transport engine.
//
// The transport package defines interfaces and implementations for:
//   - Meshtastic LoRa packet mesh (868/915 MHz)
//   - Tor v3 hidden services (.onion)
//   - I2P SAM bridge / garlic routing
//   - Bluetooth Low Energy (BLE) mesh
//
// All transport implementations must satisfy the Transport interface.
// This package MUST NOT import any Layer 2 (game engine) packages.
package transport

import (
	"context"
	"net"
	"sync"
)

// TorTransport implements the Transport interface for Tor v3 hidden services.
// It mocks the Tor hardware for development.
type TorTransport struct {
	mu      sync.Mutex
	config  Config
	inbox   chan []byte // packets received from the Tor network
	outbox  chan []byte // packets to send to the Tor network
	closeCh chan struct{}
	closed  bool
	addr    net.Addr
}

// NewTorTransport creates a new Tor transport with the given config.
// The addr parameter is the local transport address (e.g., a mock address).
func NewTorTransport(cfg Config, addr net.Addr) *TorTransport {
	tt := &TorTransport{
		config:  cfg,
		inbox:   make(chan []byte, 10),
		outbox:  make(chan []byte, 10),
		closeCh: make(chan struct{}),
		addr:    addr,
	}
	go tt.run()
	return tt
}

// run is the main goroutine that handles sending and receiving packets.
// In a real implementation, this would interface with the Tor daemon via the control port.
func (tt *TorTransport) run() {
	defer close(tt.inbox)
	defer close(tt.outbox)
	for {
		select {
		case <-tt.closeCh:
			return
		case pkt := <-tt.outbox:
			// Simulate sending packet to the Tor network.
			// For now, we immediately loop back to the inbox to simulate a reliable link.
			select {
			case tt.inbox <- pkt:
			case <-tt.closeCh:
				return
			}
		}
	}
}

// Send transmits a frame to the specified peer address.
func (tt *TorTransport) Send(ctx context.Context, addr net.Addr, frame []byte) (int, error) {
	if tt.isClosed() {
		return 0, ErrClosed
	}
	// Enforce max frame size (Tor: 4096 bytes per spec? but we use config).
	if len(frame) > tt.config.MaxFrameSize && tt.config.MaxFrameSize > 0 {
		return 0, ErrFrameTooLarge
	}
	select {
	case tt.outbox <- frame:
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	return len(frame), nil
}

// Recv receives a frame from the transport.
// It waits for a packet from the Tor network.
func (tt *TorTransport) Recv(ctx context.Context) ([]byte, net.Addr, error) {
	if tt.isClosed() {
		return nil, nil, ErrClosed
	}
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case pkt := <-tt.inbox:
		// For simplicity, we return the local address as the sender.
		// In a real implementation, the sender address would be extracted from the Tor stream metadata.
		return pkt, tt.addr, nil
	}
}

// Addr returns the local transport address.
func (tt *TorTransport) Addr() net.Addr {
	tt.mu.Lock()
	defer tt.mu.Unlock()
	return tt.addr
}

// Close gracefully shuts down the transport.
func (tt *TorTransport) Close() error {
	if tt.isClosed() {
		return nil
	}
	tt.mu.Lock()
	defer tt.mu.Unlock()
	if tt.closed {
		tt.mu.Unlock()
		return nil
	}
	tt.closed = true
	close(tt.closeCh)
	return nil
}

// String returns the transport name for logging/debugging.
func (tt *TorTransport) String() string {
	return "Tor"
}

// isClosed checks if the transport is closed.
func (tt *TorTransport) isClosed() bool {
	tt.mu.Lock()
	defer tt.mu.Unlock()
	return tt.closed
}
