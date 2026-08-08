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

// I2PTransport implements the Transport interface for I2P SAM garlic tunnels.
// It mocks the I2P hardware for development.
type I2PTransport struct {
	mu      sync.Mutex
	config  Config
	inbox   chan []byte // packets received from the I2P network
	outbox  chan []byte // packets to send to the I2P network
	closeCh chan struct{}
	closed  bool
	addr    net.Addr
}

// NewI2PTransport creates a new I2P transport with the given config.
// The addr parameter is the local transport address (e.g., a mock address).
func NewI2PTransport(cfg Config, addr net.Addr) *I2PTransport {
	it := &I2PTransport{
		config:  cfg,
		inbox:   make(chan []byte, 10),
		outbox:  make(chan []byte, 10),
		closeCh: make(chan struct{}),
		addr:    addr,
	}
	go it.run()
	return it
}

// run is the main goroutine that handles sending and receiving packets.
// In a real implementation, this would interface with the I2P SAM bridge.
func (it *I2PTransport) run() {
	defer close(it.inbox)
	defer close(it.outbox)
	for {
		select {
		case <-it.closeCh:
			return
		case pkt := <-it.outbox:
			// Simulate sending packet to the I2P network.
			// For now, we immediately loop back to the inbox to simulate a reliable link.
			select {
			case it.inbox <- pkt:
			case <-it.closeCh:
				return
			}
		}
	}
}

// Send transmits a frame to the specified peer address.
func (it *I2PTransport) Send(ctx context.Context, addr net.Addr, frame []byte) (int, error) {
	if it.isClosed() {
		return 0, ErrClosed
	}
	// Enforce max frame size (I2P: 1024 bytes per spec? but we use config).
	if len(frame) > it.config.MaxFrameSize && it.config.MaxFrameSize > 0 {
		return 0, ErrFrameTooLarge
	}
	select {
	case it.outbox <- frame:
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	return len(frame), nil
}

// Recv receives a frame from the transport.
// It waits for a packet from the I2P network.
func (it *I2PTransport) Recv(ctx context.Context) ([]byte, net.Addr, error) {
	if it.isClosed() {
		return nil, nil, ErrClosed
	}
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case pkt := <-it.inbox:
		// For simplicity, we return the local address as the sender.
		// In a real implementation, the sender address would be extracted from the I2P session metadata.
		return pkt, it.addr, nil
	}
}

// Addr returns the local transport address.
func (it *I2PTransport) Addr() net.Addr {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.addr
}

// Close gracefully shuts down the transport.
func (it *I2PTransport) Close() error {
	if it.isClosed() {
		return nil
	}
	it.mu.Lock()
	defer it.mu.Unlock()
	if it.closed {
		it.mu.Unlock()
		return nil
	}
	it.closed = true
	close(it.closeCh)
	return nil
}

// String returns the transport name for logging/debugging.
func (it *I2PTransport) String() string {
	return "I2P"
}

// isClosed checks if the transport is closed.
func (it *I2PTransport) isClosed() bool {
	it.mu.Lock()
	defer it.mu.Unlock()
	return it.closed
}
