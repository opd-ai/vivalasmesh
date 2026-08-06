// Package types provides shared type definitions for Viva Las Mesh.
package types

import "time"

// Crypto types
type (
	PublicKey  []byte
	PrivateKey []byte
	SharedKey  []byte
	Nonce      []byte
	Signature  []byte
)

// NoiseIKState represents the Noise IK handshake state.
type NoiseIKState int

const (
	NoiseIKInit NoiseIKState = iota
	NoiseIKSentE
	NoiseIKReceivedE
	NoiseIKSentS
	NoiseIKComplete
)

// CRDT types
type CRDTType string

const (
	CRDTTypeLWWRegister CRDTType = "lww_register"
	CRDTTypeORSet       CRDTType = "or_set"
	CRDTTypeCounter     CRDTType = "counter"
	CRDTTypeRegister    CRDTType = "register"
)

// CRDTOperation represents a CRDT operation.
type CRDTOperation struct {
	Type      CRDTOpType
	Key       string
	Value     []byte
	Timestamp time.Time
	NodeID    NodeID
}

// CRDTOpType identifies the CRDT operation type.
type CRDTOpType string

const (
	CRDTOpSet    CRDTOpType = "set"
	CRDTOpAdd    CRDTOpType = "add"
	CRDTOpRemove CRDTOpType = "remove"
	CRDTOpInc    CRDTOpType = "inc"
	CRDTOpDec    CRDTOpType = "dec"
)
