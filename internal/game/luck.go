// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// RatPackMember represents the members of the Rat Pack Pantheon.
type RatPackMember string

const (
	RatPackSinatra RatPackMember = "FRANK_SINATRA"
	RatPackMartin  RatPackMember = "DEAN_MARTIN"
	RatPackDavis   RatPackMember = "SAMMY_DAVIS_JR"
	RatPackHughes  RatPackMember = "HOWARD_HUGHES"
)

// LuckEntropy represents the current luck entropy state, influenced by offerings to the Rat Pack Pantheon.
// Higher luck entropy means better luck; lower means worse luck.
type LuckEntropy int

// OfferRating represents the strength of an offering to a Rat Pack member.
type OfferRating int

const (
	OfferRatingNone   OfferRating = 0 // No offering
	OfferRatingLow    OfferRating = 1 // Small offering
	OfferRatingMedium OfferRating = 2 // Moderate offering
	OfferRatingHigh   OfferRating = 3 // Generous offering
	OfferRatingMax    OfferRating = 4 // Lavish offering
)

// GetLuckFromOfferings calculates the luck entropy based on offerings made to the Rat Pack Pantheon.
// Each member contributes to luck based on the offering rating.
// The formula is a simple weighted sum.
func GetLuckFromOfferings(sinatra, martin, davis, hughes OfferRating) LuckEntropy {
	// Weights for each member (reflecting their influence on luck)
	const (
		weightSinatra = 4
		weightMartin  = 3
		weightDavis   = 2
		weightHughes  = 1
	)

	// Calculate raw luck score
	raw := sinatra*weightSinatra +
		martin*weightMartin +
		davis*weightDavis +
		hughes*weightHughes

	// Convert to LuckEntropy (we'll use a range of -20 to +20 for example)
	// Map raw score (0-4*4+4*3+4*2+4*1 = 16+12+8+4=40) to -20..20
	// Luck = (raw - 20) * 2 / 2? Let's just clamp and shift.
	// We'll map raw 0..40 to -20..20: luck = raw - 20
	luck := raw - 20
	if luck < -20 {
		luck = -20
	}
	if luck > 20 {
		luck = 20
	}
	return LuckEntropy(luck)
}

// GetLuckDescription returns a description of the current luck entropy.
func GetLuckDescription(luck LuckEntropy) string {
	switch {
	case luck <= -15:
		return "Cursed: The Rat Pack has turned their backs; misfortune follows."
	case luck <= -5:
		return "Unlucky: Minor setbacks and delays are common."
	case luck <= 5:
		return "Neutral: Luck is balanced; neither fortune nor famine."
	case luck <= 15:
		return "Fortunate: The Rat Pack smiles upon you; opportunities arise."
	default:
		return "Blessed: Lady Luck favors you greatly; success is likely."
	}
}
