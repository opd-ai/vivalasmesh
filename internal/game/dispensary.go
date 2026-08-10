// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

// DispensaryItem represents a chemical item available from dispensaries.
type DispensaryItem string

const (
	DispensaryCHDIST DispensaryItem = "CH-DIST" // THC Distillate
	DispensaryCHCART DispensaryItem = "CH-CART" // Live Resin
	DispensaryCHSODA DispensaryItem = "CH-SODA" // Nano-THC Soda
	DispensaryCHMEL  DispensaryItem = "CH-MEL"  // Melatonin
	DispensaryCHSHRM DispensaryItem = "CH-SHRM" // Psilocybin
	DispensaryCHWAX  DispensaryItem = "CH-WAX"  // THCA Diamond Wax
)

// ChemicalIntakeResult represents the effect of consuming a dispensary item.
type ChemicalIntakeResult struct {
	THCDosage          int // THC dosage units (affects Buzz and Paranoia)
	AlcoholLevel       int // Alcohol level units (affects Paranoia)
	MetabolicTicks     int // Metabolic ticks (affects hunger, fatigue?)
	HydrationChange    int // Change in hydration percentage (can be positive or negative)
	StreetSmartsChange int // Change in street smarts (could be negative for some substances)
	HealthChange       int // Change in health (could be negative for harmful substances)
	BuzzChange         int // Change in buzz (0-10)
	DignityChange      int // Change in dignity
	HardwareChange     int // Change in hardware (technical skill?)
}

// ConsumeDispensaryItem processes the intake of a dispensary item and returns the resulting
// chemical effects on the player's physiology.
// In a real implementation, this would be more detailed, with varying potency, tolerance, etc.
func ConsumeDispensaryItem(item DispensaryItem) ChemicalIntakeResult {
	switch item {
	case DispensaryCHDIST:
		// THC Distillate: high THC, low alcohol, slight hydration decrease
		return ChemicalIntakeResult{
			THCDosage:          15,
			AlcoholLevel:       0,
			MetabolicTicks:     5,
			HydrationChange:    -5,
			StreetSmartsChange: -2,
			HealthChange:       0,
			BuzzChange:         3,
			DignityChange:      -1,
			HardwareChange:     -1,
		}
	case DispensaryCHCART:
		// Live Resin: medium THC, no alcohol
		return ChemicalIntakeResult{
			THCDosage:          12,
			AlcoholLevel:       0,
			MetabolicTicks:     4,
			HydrationChange:    -3,
			StreetSmartsChange: -1,
			HealthChange:       0,
			BuzzChange:         2,
			DignityChange:      0,
			HardwareChange:     0,
		}
	case DispensaryCHSODA:
		// Nano-THC Soda: low THC, high sugar, slight hydration increase
		return ChemicalIntakeResult{
			THCDosage:          5,
			AlcoholLevel:       0,
			MetabolicTicks:     2,
			HydrationChange:    3, // Soda hydrates
			StreetSmartsChange: 0,
			HealthChange:       -1, // Sugar crash later
			BuzzChange:         1,
			DignityChange:      0,
			HardwareChange:     0,
		}
	case DispensaryCHMEL:
		// Melatonin: no THC, promotes sleep, reduces buzz
		return ChemicalIntakeResult{
			THCDosage:          0,
			AlcoholLevel:       0,
			MetabolicTicks:     1,
			HydrationChange:    0,
			StreetSmartsChange: 0,
			HealthChange:       0,
			BuzzChange:         -2, // Reduces buzz
			DignityChange:      0,
			HardwareChange:     0,
		}
	case DispensaryCHSHRM:
		// Psilocybin: psychedelic, affects perception, may increase street smarts? but can cause paranoia
		return ChemicalIntakeResult{
			THCDosage:          3, // Mild psychedelic effect
			AlcoholLevel:       0,
			MetabolicTicks:     3,
			HydrationChange:    0,
			StreetSmartsChange: 2, // May increase perception
			HealthChange:       0,
			BuzzChange:         1,
			DignityChange:      -2, // May cause erratic behavior
			HardwareChange:     -2, // May impair technical skill
		}
	case DispensaryCHWAX:
		// THCA Diamond Wax: very high THC, potent
		return ChemicalIntakeResult{
			THCDosage:          20,
			AlcoholLevel:       0,
			MetabolicTicks:     8,
			HydrationChange:    -10,
			StreetSmartsChange: -3,
			HealthChange:       0,
			BuzzChange:         4,
			DignityChange:      -2,
			HardwareChange:     -2,
		}
	default:
		return ChemicalIntakeResult{}
	}
}
