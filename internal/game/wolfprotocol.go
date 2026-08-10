// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// WolfProtocolMethod represents the different body disposal methods available through "The Wolf Protocol".
type WolfProtocolMethod string

const (
	WolfProtocolDPACID WolfProtocolMethod = "DP-ACID"  // Acid Drains
	WolfProtocolDPLAKE WolfProtocolMethod = "DP-LAKE"  // Hoover Dam
	WolfProtocolDPBURN WolfProtocolMethod = "DP-BURN"  // Incinerator
	WolfProtocolDPRUNK WolfProtocolMethod = "DP-TRUNK" // Rival Parking
)

// WolfProtocolResult represents the outcome of using a Wolf Protocol disposal method.
type WolfProtocolResult struct {
	Success        bool           // Whether the disposal was successful
	HeatChange     int            // Change in police heat (positive increases heat, negative decreases)
	FactionChanges map[string]int // Changes in faction standing (e.g., "aec_rep": -10)
	CreditsChange  int            // Change in credits (could be negative if disposal costs money)
	Description    string         // Description of the outcome
}

// UseWolfProtocol disposes of a body using the specified Wolf Protocol method.
// Returns the result of the disposal attempt.
// In a real implementation, this would involve skill checks, risks, and rewards.
func UseWolfProtocol(method WolfProtocolMethod) WolfProtocolResult {
	switch method {
	case WolfProtocolDPACID:
		// Acid Drains: effective but risky; may attract attention if not done properly
		return WolfProtocolResult{
			Success:    true,
			HeatChange: 5, // Moderate heat increase
			FactionChanges: map[string]int{
				"metro_pd_rep": -10, // Police dislike evidence destruction
				"darknet_rep":  5,   // Darknet hackers appreciate discreet disposal
			},
			CreditsChange: 0,
			Description:   "Body dissolved in acid drain. Minimal traces left.",
		}
	case WolfProtocolDPLAKE:
		// Hoover Dam: deep water disposal, very effective but requires transport
		return WolfProtocolResult{
			Success:    true,
			HeatChange: 2, // Low heat increase (if done discreetly)
			FactionChanges: map[string]int{
				"metro_pd_rep":  -5,
				"corporate_rep": -5, // Corporations dislike environmental damage
			},
			CreditsChange: -20, // Cost of transport and bribes
			Description:   "Body weighted and sunk in Hoover Dam. Nearly impossible to recover.",
		}
	case WolfProtocolDPBURN:
		// Incinerator: high heat, destroys evidence completely but may leave traces of smoke
		return WolfProtocolResult{
			Success:    true,
			HeatChange: 8, // High heat increase due to smoke and energy signature
			FactionChanges: map[string]int{
				"metro_pd_rep":  -15, // Police notice unusual incinerator activity
				"corporate_rep": 10,  // Corporations appreciate efficient waste disposal
			},
			CreditsChange: -10, // Cost of incinerator fees
			Description:   "Body incinerated at high temperature. Ashes scattered.",
		}
	case WolfProtocolDPRUNK:
		// Rival Parking: dumping in a rival's territory, risky but can frame others
		return WolfProtocolResult{
			Success:    true,
			HeatChange: 3, // Moderate heat increase
			FactionChanges: map[string]int{
				"chicago_outfit_rep": -20, // Rival gang very angry if traced
				"darknet_rep":        10,  // Darknet may profit from the chaos
			},
			CreditsChange: 0,
			Description:   "Body dumped in rival parking lot. High risk of gang retaliation.",
		}
	default:
		return WolfProtocolResult{
			Success:        false,
			HeatChange:     0,
			FactionChanges: nil,
			CreditsChange:  0,
			Description:    "Unknown disposal method.",
		}
	}
}
