// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"github.com/opd-ai/vivalasmesh/pkg/types"
)

// AnachronismFlag represents different types of anachronisms that can be detected.
type AnachronismFlag string

const (
	FlagAlienTech AnachronismFlag = "FLAG_ALIEN_TECH" // Future tech in past era
	FlagWireBug   AnachronismFlag = "FLAG_WIRE_BUG"   // Surveillance tech in inappropriate era
	FlagAntique   AnachronismFlag = "FLAG_ANTIQUE"    // Past tech in future era
)

// DetectAnachronism checks if an item is anachronistic in the current era.
// Returns the appropriate flag if anachronistic, or empty string if not.
func DetectAnachronism(itemEra, currentEra types.Era) AnachronismFlag {
	// Define what constitutes an anachronism based on item era vs current era
	switch {
	case itemEra == types.Era2026Cyber && currentEra == types.Era1953Atomic:
		return FlagAlienTech
	case itemEra == types.Era2026Cyber && currentEra == types.Era1962Syndicate:
		return FlagWireBug
	case itemEra == types.Era1953Atomic && currentEra == types.Era2026Cyber:
		return FlagAntique
	// Add more combinations as needed
	default:
		return "" // Not anachronistic
	}
}

// GetAnachronismDetails returns the penalty details for a given anachronism flag.
// Returns: cinematic reaction description, faction standing changes
func GetAnachronismDetails(flag AnachronismFlag) (string, map[string]int) {
	switch flag {
	case FlagAlienTech:
		// AEC Agents lock down sector; suspect Soviet/UFO spy. AEC Rep -50
		return "AEC Agents lock down sector; suspect Soviet/UFO spy.", map[string]int{
			"aec_rep": -50,
		}
	case FlagWireBug:
		// Mob Capos suspect wiretap device; draw snubnose .38s. Outfit Rep -40
		return "Mob Capos suspect wiretap device; draw snubnose .38s.", map[string]int{
			"outfit_rep": -40,
		}
	case FlagAntique:
		// Pawn Brokers offer double cash; street hackers laugh. No Hit (+10 Cred)
		return "Pawn Brokers offer double cash; street hackers laugh.", map[string]int{
			"credits": 10, // Bonus credits
		}
	default:
		return "", nil
	}
}
