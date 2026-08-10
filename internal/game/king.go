// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"github.com/opd-ai/vivalasmesh/pkg/types"
)

// KingManifestation represents the different forms of The Transcendent King (Elvis).
type KingManifestation string

const (
	ManifestationProphet   KingManifestation = "PROPHET"   // ERA 1: Sun Records Rockabilly
	ManifestationHollywood KingManifestation = "HOLLYWOOD" // ERA 2: Hollywood Gold-Suit
	ManifestationOracle    KingManifestation = "ORACLE"    // ERA 3: White Jumpsuit Oracle
	ManifestationCodeGod   KingManifestation = "CODE_GOD"  // ERA 4: Holographic Spectral Code-God
)

// GetCurrentManifestation returns the current manifestation of The King based on era.
func GetCurrentManifestation(currentEra types.Era) KingManifestation {
	switch currentEra {
	case types.Era1953Atomic:
		return ManifestationProphet
	case types.Era1962Syndicate:
		return ManifestationHollywood
	case types.Era1993Corporate:
		return ManifestationOracle
	case types.Era2026Cyber:
		return ManifestationCodeGod
	default:
		return ManifestationProphet // Default to prophet
	}
}

// ApplyKingEffect applies the effect of the current King manifestation.
// In a full implementation, this would modify game state, attributes, etc.
// For now, it returns a description of what the effect does.
func ApplyKingEffect(currentEra types.Era) string {
	switch GetCurrentManifestation(currentEra) {
	case ManifestationProphet:
		// ERA 1: Prophet form → Prophetic Riff: +4 Hustle, immunity to fear
		return "Prophetic Riff: Grants +4 Hustle and permanent immunity to fear."
	case ManifestationHollywood:
		// ERA 2: Hollywood Gold-Suit → Velvet Charm: +5 Dignity, Mob Capos refuse to attack
		return "Velvet Charm: Grants +5 Dignity; Mob Capos refuse to attack player."
	case ManifestationOracle:
		// ERA 3: White Jumpsuit Oracle → Neon Redemption: Clears all Paranoia, 1-time Revive token
		return "Neon Redemption: Clears all Paranoia; grants 1-time Revive token."
	case ManifestationCodeGod:
		// ERA 4: Holographic Spectral Code-God → Graceland Protocol: Overrides all local casino SCADA systems for 30s
		return "Graceland Protocol: Overrides all local casino SCADA systems for 30s."
	default:
		return "Unknown King manifestation."
	}
}
