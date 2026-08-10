// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// ScarType represents the different types of permanent scarring that can occur.
type ScarType string

const (
	ScarTRNOSE   ScarType = "TR-NOSE"   // Nasal scarring
	ScarTRFINGER ScarType = "TR-FINGER" // Finger scarring
	ScarTRBURN   ScarType = "TR-BURN"   // Burn scarring
	ScarTRIMP    ScarType = "TR-LIMP"   // Limping due to leg/foot injury
)

// ScarModifier represents the permanent attribute modifications caused by a scar.
type ScarModifier struct {
	HealthChange       int // Permanent change to health max or current?
	BuzzChange         int // Permanent change to buzz tolerance?
	HustleChange       int // Permanent change to hustle (speed, agility)?
	StreetSmartsChange int // Permanent change to street smarts?
	DignityChange      int // Permanent change to dignity (scars may affect social interactions)
	HardwareChange     int // Permanent change to hardware (technical skill, fine motor control)?
}

// GetScarModifier returns the attribute modifiers for the given scar type.
// These are permanent reductions (or rarely increases) to attributes.
func GetScarModifier(scar ScarType) ScarModifier {
	switch scar {
	case ScarTRNOSE:
		// Nasal scarring: reduced breathing capacity, affects hustle and health
		return ScarModifier{
			HealthChange:       -5,
			BuzzChange:         0,
			HustleChange:       -3,
			StreetSmartsChange: 0,
			DignityChange:      -2,
			HardwareChange:     0,
		}
	case ScarTRFINGER:
		// Finger scarring: reduced dexterity, affects hardware and street smarts (lockpicking, etc.)
		return ScarModifier{
			HealthChange:       0,
			BuzzChange:         0,
			HustleChange:       0,
			StreetSmartsChange: -2,
			DignityChange:      -1,
			HardwareChange:     -3,
		}
	case ScarTRBURN:
		// Burn scarring: painful, affects health, buzz, dignity
		return ScarModifier{
			HealthChange:       -8,
			BuzzChange:         -2, // Pain reduces buzz tolerance?
			HustleChange:       -4,
			StreetSmartsChange: -1,
			DignityChange:      -3,
			HardwareChange:     -2,
		}
	case ScarTRIMP:
		// Limping: reduces movement speed, affects hustle
		return ScarModifier{
			HealthChange:       -3,
			BuzzChange:         0,
			HustleChange:       -6, // Significant reduction in speed/agility
			StreetSmartsChange: 0,
			DignityChange:      -1,
			HardwareChange:     0,
		}
	default:
		return ScarModifier{}
	}
}
