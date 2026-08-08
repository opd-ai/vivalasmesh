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
	"errors"
	"net"
	"sync"

	"github.com/fxamacker/cbor/v2"
)

// LoRaTransport implements the Transport interface for Meshtastic LoRa radios.
// It uses CBOR for packet serialization and mocks the radio hardware for development.
type LoRaTransport struct {
	mu      sync.Mutex
	config  Config
	inbox   chan []byte // encoded CBOR packets received from the radio
	outbox  chan []byte // encoded CBOR packets to send to the radio
	closeCh chan struct{}
	closed  bool
	addr    net.Addr
}

// NewLoRaTransport creates a new LoRa transport with the given config.
// The addr parameter is the local transport address (e.g., a mock address).
func NewLoRaTransport(cfg Config, addr net.Addr) *LoRaTransport {
	lt := &LoRaTransport{
		config:  cfg,
		inbox:   make(chan []byte, 10),
		outbox:  make(chan []byte, 10),
		closeCh: make(chan struct{}),
		addr:    addr,
	}
	go lt.run()
	return lt
}

// run is the main goroutine that handles sending and receiving packets.
// In a real implementation, this would interface with the radio hardware via UART/SPI.
func (lt *LoRaTransport) run() {
	defer close(lt.inbox)
	defer close(lt.outbox)
	for {
		select {
		case <-lt.closeCh:
			return
		case encoded := <-lt.outbox:
			// Simulate sending encoded packet to the radio.
			// For now, we immediately loop back to the inbox to simulate a reliable link.
			// In reality, this would go out the radio port and incoming packets would come from another source.
			select {
			case lt.inbox <- encoded:
			case <-lt.closeCh:
				return
			}
		}
	}
}

// Send transmits a frame to the specified peer address.
// The frame is CBOR-encoded before transmission.
func (lt *LoRaTransport) Send(ctx context.Context, addr net.Addr, frame []byte) (int, error) {
	if lt.isClosed() {
		return 0, ErrClosed
	}
	// CBOR-encode the frame.
	encoded, err := cbor.Marshal(frame)
	if err != nil {
		return 0, err
	}
	// Enforce max frame size (Meshtastic LoRa: 237 bytes).
	if len(encoded) > lt.config.MaxFrameSize && lt.config.MaxFrameSize > 0 {
		return 0, ErrFrameTooLarge
	}
	select {
	case lt.outbox <- encoded:
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	return len(encoded), nil
}

// Recv receives a frame from the transport.
// It waits for a CBOR-encoded packet from the radio, decodes it, and returns the frame.
func (lt *LoRaTransport) Recv(ctx context.Context) ([]byte, net.Addr, error) {
	if lt.isClosed() {
		return nil, nil, ErrClosed
	}
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case encoded := <-lt.inbox:
		var frame []byte
		if err := cbor.Unmarshal(encoded, &frame); err != nil {
			return nil, nil, err
		}
		// For simplicity, we return the local address as the sender.
		// In a real implementation, the sender address would be extracted from the packet metadata.
		return frame, lt.addr, nil
	}
}

// Addr returns the local transport address.
func (lt *LoRaTransport) Addr() net.Addr {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	return lt.addr
}

// Close gracefully shuts down the transport.
func (lt *LoRaTransport) Close() error {
	if lt.isClosed() {
		return nil
	}
	lt.mu.Lock()
	defer lt.mu.Unlock()
	if lt.closed {
		lt.mu.Unlock()
		return nil
	}
	lt.closed = true
	close(lt.closeCh)
	return nil
}

// String returns the transport name for logging/debugging.
func (lt *LoRaTransport) String() string {
	return "LoRa"
}

// isClosed checks if the transport is closed.
func (lt *LoRaTransport) isClosed() bool {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	return lt.closed
}

// Errors.
var (
	ErrClosed        = errors.New("transport closed")
	ErrFrameTooLarge = errors.New("frame exceeds maximum size")
)
