// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package types provides shared type definitions for Viva Las Mesh.
package types

// Layer 2: Game Mechanics Types

// Era represents a temporal era.
type Era int

const (
	Era1953Atomic    Era = 1953 // Atomic AEC Era
	Era1962Syndicate Era = 1962 // Syndicate/Stardust Era
	Era1993Corporate Era = 1993 // Corporate Mega-Resort Era
	Era2026Cyber     Era = 2026 // Cyber-Strip Dystopia
)

// String returns the era name.
func (e Era) String() string {
	switch e {
	case Era1953Atomic:
		return "1953: Atomic AEC Era"
	case Era1962Syndicate:
		return "1962: Syndicate/Stardust Era"
	case Era1993Corporate:
		return "1993: Corporate Mega-Resort Era"
	case Era2026Cyber:
		return "2026: Cyber-Strip Dystopia"
	default:
		return "Unknown Era"
	}
}

// Position represents a world position.
type Position struct {
	X, Y int
	Era  Era
}

// ActionPriority represents action priority tiers (T1-T5).
type ActionPriority int

const (
	PriorityDivine     ActionPriority = 1 // T1: Divine/Spectral King
	PriorityCritical   ActionPriority = 2 // T2: Critical system
	PriorityStandard   ActionPriority = 3 // T3: Standard actions
	PriorityHeavy      ActionPriority = 4 // T4: Heavy actions
	PriorityBackground ActionPriority = 5 // T5: Background/environmental
)

// AttributeType represents a character attribute.
type AttributeType string

const (
	AttrHustle       AttributeType = "hustle"
	AttrHardware     AttributeType = "hardware"
	AttrStreetSmarts AttributeType = "street_smarts"
	AttrDignity      AttributeType = "dignity"
	AttrLuck         AttributeType = "luck"
	AttrParanoia     AttributeType = "paranoia"
)

// ItemType represents an item category.
type ItemType string

const (
	ItemWeapon     ItemType = "weapon"
	ItemCybernetic ItemType = "cybernetic"
	ItemConsumable ItemType = "consumable"
	ItemKey        ItemType = "key"
	ItemTool       ItemType = "tool"
	ItemData       ItemType = "data"
	ItemCurrency   ItemType = "currency"
)

// FactionType represents a faction.
type FactionType string

const (
	FactionPlayer    FactionType = "player"
	FactionMetroPD   FactionType = "metro_pd"
	FactionSyndicate FactionType = "syndicate"
	FactionCorporate FactionType = "corporate"
	FactionAnarchist FactionType = "anarchist"
	FactionCult      FactionType = "cult"
	FactionSpectral  FactionType = "spectral"
	FactionNeutral   FactionType = "neutral"
)

// NPCArchetype represents an NPC archetype.
type NPCArchetype string

const (
	NPCDealer       NPCArchetype = "dealer"
	NPCFixer        NPCArchetype = "fixer"
	NPCHacker       NPCArchetype = "hacker"
	NPCEnforcer     NPCArchetype = "enforcer"
	NPCCourtesan    NPCArchetype = "courtesan"
	NPCOfficer      NPCArchetype = "officer"
	NPCScientist    NPCArchetype = "scientist"
	NPCDrifter      NPCArchetype = "drifter"
	NPCSpectralKing NPCArchetype = "spectral_king"
	NPCCooler       NPCArchetype = "cooler"
)
