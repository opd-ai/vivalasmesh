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
	"io"
	"net"
)

// Transport defines the interface for all P2P transport implementations.
// Implementations handle frame-level send/receive with transport-specific framing.
// All methods are safe for concurrent use.
type Transport interface {
	// Send transmits a frame to the specified peer address.
	// Returns the number of bytes written or an error.
	Send(ctx context.Context, addr net.Addr, frame []byte) (int, error)

	// Recv receives a frame from the transport.
	// Returns the frame data, sender address, or an error.
	// Blocks until a frame is available or context is cancelled.
	Recv(ctx context.Context) ([]byte, net.Addr, error)

	// Addr returns the local transport address.
	Addr() net.Addr

	// Close gracefully shuts down the transport.
	Close() error

	// String returns the transport name for logging/debugging.
	String() string
}

// Frame represents a transport frame with metadata.
type Frame struct {
	// Data is the frame payload.
	Data []byte
	// Src is the source address.
	Src net.Addr
	// Dst is the destination address.
	Dst net.Addr
	// Timestamp is when the frame was received/sent.
	Timestamp int64
}

// FrameReader provides a streaming frame reader interface.
type FrameReader interface {
	ReadFrame() (*Frame, error)
	io.Closer
}

// FrameWriter provides a streaming frame writer interface.
type FrameWriter interface {
	WriteFrame(*Frame) error
	io.Closer
}

// Config holds common transport configuration.
type Config struct {
	// MaxFrameSize is the maximum frame size in bytes.
	// Meshtastic LoRa: 237 bytes
	MaxFrameSize int
	// EnableEncryption enables Noise IK framing.
	EnableEncryption bool
	// PrivateKey is the transport's static private key (Noise IK).
	PrivateKey []byte
}
