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
	"crypto/cipher"
	"crypto/rand"
	"io"
	"net"
	"sync"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

// NoiseTransport wraps a Transport with Noise IK handshake and Double Ratchet
// session management using chacha20poly1305 and curve25519.
// This is a simplified mock implementation for development.
type NoiseTransport struct {
	mu         sync.Mutex
	underlying Transport
	localKey   [32]byte // local private key
	remoteKey  [32]byte // remote public key (static for demo)
	cipher     cipher.AEAD
	hasKey     bool
	closeCh    chan struct{}
	closed     bool
}

// NewNoiseTransport wraps the given transport with Noise encryption.
// It performs a static Noise IK handshake (simplified) and sets up
// a chacha20poly1305 cipher for encrypting frames.
func NewNoiseTransport(underlying Transport) (*NoiseTransport, error) {
	// Generate a local static key pair (simplified: we only need private key for Noise IK).
	// In a real Noise IK handshake, we would exchange public keys.
	var localPriv [32]byte
	if _, err := io.ReadFull(rand.Reader, localPriv[:]); err != nil {
		return nil, err
	}
	// For demo, we assume a fixed remote public key (all zeros).
	var remotePub [32]byte
	// Compute the shared secret using curve25519.
	var sharedSecret [32]byte
	curve25519.ScalarMult(&sharedSecret, &localPriv, &remotePub)
	// Derive a key for chacha20poly1305 from the shared secret (simplified: use first 32 bytes).
	var key [32]byte
	copy(key[:], sharedSecret[:])
	cipher, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, err
	}
	nt := &NoiseTransport{
		underlying: underlying,
		localKey:   localPriv,
		remoteKey:  remotePub,
		cipher:     cipher,
		closeCh:    make(chan struct{}),
	}
	go nt.run()
	return nt, nil
}

// run is a placeholder for background tasks (e.g., ratchet stepping).
func (nt *NoiseTransport) run() {
	// In a full implementation, we would rotate keys and handle rekeying.
	<-nt.closeCh
}

// Send encrypts the frame with chacha20poly1305 before sending to the underlying transport.
func (nt *NoiseTransport) Send(ctx context.Context, addr net.Addr, frame []byte) (int, error) {
	if nt.isClosed() {
		return 0, ErrClosed
	}
	nt.mu.Lock()
	defer nt.mu.Unlock()
	if !nt.hasKey {
		// In a full implementation, we would perform the handshake here.
		// For demo, we assume the key is already agreed upon.
		nt.hasKey = true
	}
	// Create a nonce (for simplicity, we use a zero nonce; in practice, use a counter).
	nonce := make([]byte, nt.cipher.NonceSize())
	// Encrypt the frame.
	ciphertext := nt.cipher.Seal(nil, nonce, frame, nil)
	// Send the ciphertext to the underlying transport.
	return nt.underlying.Send(ctx, addr, ciphertext)
}

// Recv receives a frame from the underlying transport and decrypts it.
func (nt *NoiseTransport) Recv(ctx context.Context) ([]byte, net.Addr, error) {
	if nt.isClosed() {
		return nil, nil, ErrClosed
	}
	// Receive the ciphertext from the underlying transport.
	ciphertext, addr, err := nt.underlying.Recv(ctx)
	if err != nil {
		return nil, nil, err
	}
	nt.mu.Lock()
	defer nt.mu.Unlock()
	if !nt.hasKey {
		// Perform handshake if needed.
		nt.hasKey = true
	}
	// Create a nonce (must match the one used in Seal).
	nonce := make([]byte, nt.cipher.NonceSize())
	// Decrypt the frame.
	plaintext, err := nt.cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, nil, err
	}
	return plaintext, addr, nil
}

// Addr returns the local transport address (delegates to underlying).
func (nt *NoiseTransport) Addr() net.Addr {
	return nt.underlying.Addr()
}

// Close closes the underlying transport and signals the run goroutine to exit.
func (nt *NoiseTransport) Close() error {
	if nt.isClosed() {
		return nil
	}
	nt.mu.Lock()
	if nt.closed {
		nt.mu.Unlock()
		return nil
	}
	nt.closed = true
	close(nt.closeCh)
	nt.mu.Unlock()
	return nt.underlying.Close()
}

// String returns the transport name for logging/debugging.
func (nt *NoiseTransport) String() string {
	return "noise(" + nt.underlying.String() + ")"
}

// isClosed checks if the transport is closed.
func (nt *NoiseTransport) isClosed() bool {
	nt.mu.Lock()
	defer nt.mu.Unlock()
	return nt.closed
}
