// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package types provides shared type definitions for Viva Las Mesh.
// This package is carefully segregated by architectural layer:
//   - Layer 1 types (transport, crypto, sync, daemon) in types/layer1/
//   - Layer 2 types (engine, render, game, tui) in types/layer2/
//   - Shared/base types in types/
//
// Direct cross-layer imports are forbidden. Use interfaces defined in this package.
package types

import (
	"net"
	"time"
)

// Layer 1: P2P Infrastructure Types

// NodeID uniquely identifies a node in the mesh.
type NodeID string

// PeerID uniquely identifies a peer connection.
type PeerID string

// TransportType identifies the transport protocol.
type TransportType string

const (
	TransportLoRa TransportType = "lora"
	TransportTor  TransportType = "tor"
	TransportI2P  TransportType = "i2p"
	TransportBLE  TransportType = "ble"
)

// TransportAddr represents a transport-specific address.
type TransportAddr struct {
	Type      TransportType
	Address   string
	PublicKey []byte
}

// NetworkAddr returns the address as a net.Addr.
func (t TransportAddr) NetworkAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP(t.Address), Port: 0}
}

// Frame represents a network frame.
type Frame struct {
	ID        string
	Src       NodeID
	Dst       NodeID
	Type      FrameType
	Payload   []byte
	Timestamp time.Time
	Signature []byte
}

// FrameType identifies the frame type.
type FrameType string

const (
	FrameData      FrameType = "data"
	FrameControl   FrameType = "control"
	FrameHandshake FrameType = "handshake"
	FrameHeartbeat FrameType = "heartbeat"
	FrameSync      FrameType = "sync"
)
