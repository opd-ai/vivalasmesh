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

// BLETransport implements the Transport interface for Bluetooth Low Energy mesh.
// It mocks the BLE hardware for development.
type BLETransport struct {
	mu      sync.Mutex
	config  Config
	inbox   chan []byte // packets received from the BLE network
	outbox  chan []byte // packets to send to the BLE network
	closeCh chan struct{}
	closed  bool
	addr    net.Addr
}

// NewBLETransport creates a new BLE transport with the given config.
// The addr parameter is the local transport address (e.g., a mock address).
func NewBLETransport(cfg Config, addr net.Addr) *BLETransport {
	bt := &BLETransport{
		config:  cfg,
		inbox:   make(chan []byte, 10),
		outbox:  make(chan []byte, 10),
		closeCh: make(chan struct{}),
		addr:    addr,
	}
	go bt.run()
	return bt
}

// run is the main goroutine that handles sending and receiving packets.
// In a real implementation, this would interface with the BLE adapter via platform-specific APIs.
func (bt *BLETransport) run() {
	defer close(bt.inbox)
	defer close(bt.outbox)
	for {
		select {
		case <-bt.closeCh:
			return
		case pkt := <-bt.outbox:
			// Simulate sending packet to the BLE network.
			// For now, we immediately loop back to the inbox to simulate a reliable link.
			select {
			case bt.inbox <- pkt:
			case <-bt.closeCh:
				return
			}
		}
	}
}

// Send transmits a frame to the specified peer address.
func (bt *BLETransport) Send(ctx context.Context, addr net.Addr, frame []byte) (int, error) {
	if bt.isClosed() {
		return 0, ErrClosed
	}
	// Enforce max frame size (BLE: 23 bytes? but we use config).
	if len(frame) > bt.config.MaxFrameSize && bt.config.MaxFrameSize > 0 {
		return 0, ErrFrameTooLarge
	}
	select {
	case bt.outbox <- frame:
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	return len(frame), nil
}

// Recv receives a frame from the transport.
// It waits for a packet from the BLE network.
func (bt *BLETransport) Recv(ctx context.Context) ([]byte, net.Addr, error) {
	if bt.isClosed() {
		return nil, nil, ErrClosed
	}
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case pkt := <-bt.inbox:
		// For simplicity, we return the local address as the sender.
		// In a real implementation, the sender address would be extracted from the BLE advertisement metadata.
		return pkt, bt.addr, nil
	}
}

// Addr returns the local transport address.
func (bt *BLETransport) Addr() net.Addr {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	return bt.addr
}

// Close gracefully shuts down the transport.
func (bt *BLETransport) Close() error {
	if bt.isClosed() {
		return nil
	}
	bt.mu.Lock()
	defer bt.mu.Unlock()
	if bt.closed {
		bt.mu.Unlock()
		return nil
	}
	bt.closed = true
	close(bt.closeCh)
	return nil
}

// String returns the transport name for logging/debugging.
func (bt *BLETransport) String() string {
	return "BLE"
}

// isClosed checks if the transport is closed.
func (bt *BLETransport) isClosed() bool {
	bt.mu.Lock()
	defer bt.mu.Unlock()
	return bt.closed
}
