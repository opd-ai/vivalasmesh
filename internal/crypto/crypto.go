// Package crypto provides cryptographic primitives for Viva Las Mesh Layer 1.
// This is Layer 1 (P2P Infrastructure) - Noise IK, Double Ratchet, key management.
//
// Implements:
//   - Noise IK handshake (X25519, ChaCha20Poly1305)
//   - Signal Double Ratchet for session encryption
//   - Key derivation & rotation utilities
//   - Frame encryption/decryption middleware
//
// This package MUST NOT import any Layer 2 (game engine) packages.
package crypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/sha3"
)

// Constants for cryptographic parameters.
const (
	// KeySize is the size of X25519/Ed25519 keys (32 bytes).
	KeySize = 32
	// NonceSize is the size of ChaCha20Poly1305 nonces (12 bytes).
	NonceSize = 12
	// TagSize is the authentication tag size (16 bytes).
	TagSize = 16
	// Overhead is the total encryption overhead (nonce + tag).
	Overhead = NonceSize + TagSize
	// HashSize is the SHA3-256 output size.
	HashSize = 32
)

// Common errors.
var (
	ErrInvalidKeySize      = errors.New("invalid key size")
	ErrInvalidNonceSize    = errors.New("invalid nonce size")
	ErrDecryptionFailed    = errors.New("decryption failed: authentication tag mismatch")
	ErrKeyDerivationFailed = errors.New("key derivation failed")
)

// KeyPair represents an X25519 key pair.
type KeyPair struct {
	Private []byte
	Public  []byte
}

// GenerateKeyPair generates a new X25519 key pair.
func GenerateKeyPair() (*KeyPair, error) {
	private := make([]byte, KeySize)
	if _, err := rand.Read(private); err != nil {
		return nil, err
	}

	public, err := curve25519.X25519(private, curve25519.Basepoint)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		Private: private,
		Public:  public,
	}, nil
}

// SharedSecret computes the shared secret using X25519.
// privateKey is our private key, peerPublic is the peer's public key.
func SharedSecret(privateKey, peerPublic []byte) ([]byte, error) {
	if len(privateKey) != KeySize || len(peerPublic) != KeySize {
		return nil, ErrInvalidKeySize
	}
	return curve25519.X25519(privateKey, peerPublic)
}

// HKDF derives keys using HKDF-SHA3-256.
// ikm is the input key material, salt and info are optional.
func HKDF(ikm, salt, info []byte, length int) ([]byte, error) {
	hkdf := hkdf.New(sha3.New256, ikm, salt, info)
	key := make([]byte, length)
	if _, err := hkdf.Read(key); err != nil {
		return nil, ErrKeyDerivationFailed
	}
	return key, nil
}

// Encrypt encrypts plaintext using ChaCha20Poly1305.
// key must be 32 bytes, nonce must be 12 bytes.
// Returns nonce || ciphertext || tag.
func Encrypt(key, nonce, plaintext, aad []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrInvalidKeySize
	}
	if len(nonce) != NonceSize {
		return nil, ErrInvalidNonceSize
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	ciphertext := aead.Seal(nil, nonce, plaintext, aad)
	// Prepend nonce for storage/transmission
	result := make([]byte, NonceSize+len(ciphertext))
	copy(result[:NonceSize], nonce)
	copy(result[NonceSize:], ciphertext)
	return result, nil
}

// Decrypt decrypts ciphertext using ChaCha20Poly1305.
// key must be 32 bytes. ciphertext must have nonce prepended (first 12 bytes).
func Decrypt(key, ciphertext, aad []byte) ([]byte, error) {
	if len(key) != KeySize {
		return nil, ErrInvalidKeySize
	}
	if len(ciphertext) < NonceSize+TagSize {
		return nil, ErrInvalidNonceSize
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonce := ciphertext[:NonceSize]
	ct := ciphertext[NonceSize:]
	plaintext, err := aead.Open(nil, nonce, ct, aad)
	if err != nil {
		return nil, ErrDecryptionFailed
	}
	return plaintext, nil
}

// NoiseIKHandshake represents the Noise IK handshake state.
// This is a minimal implementation - production would use a full Noise library.
type NoiseIKHandshake struct {
	staticKey    *KeyPair
	ephemeralKey *KeyPair
	remoteStatic []byte
	ck           []byte // Chaining key
	h            []byte // Handshake hash
	initiator    bool
}

// NewNoiseIKInitiator creates a new Noise IK initiator handshake.
func NewNoiseIKInitiator(staticKey *KeyPair) *NoiseIKHandshake {
	return &NoiseIKHandshake{
		staticKey: staticKey,
		initiator: true,
	}
}

// NewNoiseIKResponder creates a new Noise IK responder handshake.
func NewNoiseIKResponder(staticKey *KeyPair) *NoiseIKHandshake {
	return &NoiseIKHandshake{
		staticKey: staticKey,
		initiator: false,
	}
}

// MixKey mixes a new key into the chaining key using HKDF.
// Returns the new chaining key and the derived cipher key.
func (n *NoiseIKHandshake) MixKey(inputKey []byte) ([]byte, []byte) {
	ck := HKDFOrPanic(n.ck, inputKey, []byte{}, HashSize*2)
	n.ck = ck[:HashSize]
	return n.ck, ck[HashSize:]
}

// MixHash mixes data into the handshake hash.
func (n *NoiseIKHandshake) MixHash(data []byte) {
	h := sha3.New256()
	h.Write(n.h)
	h.Write(data)
	n.h = h.Sum(nil)
}

// HKDFOrPanic is a helper that panics on error (for internal use only).
func HKDFOrPanic(ikm, salt, info []byte, length int) []byte {
	key, err := HKDF(ikm, salt, info, length)
	if err != nil {
		panic(err)
	}
	return key
}
