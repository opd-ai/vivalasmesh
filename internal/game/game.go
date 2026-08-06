// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
//
// Implements:
//   - Attribute system (Hustle, Hardware, Street Smarts, Dignity, etc.)
//   - Inventory/Items system
//   - Character/NPC system
//   - Heist/Mission system
//   - World Map/Grid with A* pathfinding
//   - FOV/Lighting system
//
// This package MUST NOT import any Layer 1 (transport/crypto/sync/daemon) packages.
package game

import (
	"errors"
)

// Common errors.
var (
	ErrInvalidAttribute = errors.New("invalid attribute")
	ErrAttributeTooLow  = errors.New("attribute too low for action")
	ErrInventoryFull    = errors.New("inventory full")
	ErrItemNotFound     = errors.New("item not found")
	ErrInvalidEntity    = errors.New("invalid entity")
)

// Attribute represents a character attribute.
type Attribute string

const (
	AttrHustle       Attribute = "hustle"        // Speed, reflexes, initiative
	AttrHardware     Attribute = "hardware"      // Tech, hacking, cybernetics
	AttrStreetSmarts Attribute = "street_smarts" // Social, perception, trade
	AttrDignity      Attribute = "dignity"       // Willpower, resistance, presence
	AttrLuck         Attribute = "luck"          // Fortune, entropy, jackpot
	AttrParanoia     Attribute = "paranoia"      // Awareness, detection, Bat Country
)

// Attributes holds all character attributes.
type Attributes struct {
	values map[Attribute]int
}

// NewAttributes creates a new attribute set with default values.
func NewAttributes() *Attributes {
	return &Attributes{
		values: map[Attribute]int{
			AttrHustle:       10,
			AttrHardware:     10,
			AttrStreetSmarts: 10,
			AttrDignity:      10,
			AttrLuck:         10,
			AttrParanoia:     0,
		},
	}
}

// Get returns an attribute value.
func (a *Attributes) Get(attr Attribute) (int, error) {
	v, ok := a.values[attr]
	if !ok {
		return 0, ErrInvalidAttribute
	}
	return v, nil
}

// Set sets an attribute value (clamped to 0-20).
func (a *Attributes) Set(attr Attribute, value int) error {
	if _, ok := a.values[attr]; !ok {
		return ErrInvalidAttribute
	}
	if value < 0 {
		value = 0
	}
	if value > 20 {
		value = 20
	}
	a.values[attr] = value
	return nil
}

// Mod modifies an attribute by delta.
func (a *Attributes) Mod(attr Attribute, delta int) error {
	v, err := a.Get(attr)
	if err != nil {
		return err
	}
	return a.Set(attr, v+delta)
}

// Check rolls an attribute check (d20 + attribute vs DC).
// Returns success, margin of success/failure.
func (a *Attributes) Check(attr Attribute, dc int) (bool, int) {
	v, _ := a.Get(attr)
	roll := 10 // In real impl, this would be random 1-20
	total := roll + v
	return total >= dc, total - dc
}

// All returns a copy of all attributes.
func (a *Attributes) All() map[Attribute]int {
	result := make(map[Attribute]int)
	for k, v := range a.values {
		result[k] = v
	}
	return result
}

// Entity represents a game entity (player, NPC, item, etc.).
type Entity struct {
	ID         string
	Type       string
	Position   Position
	Attributes *Attributes
	Faction    string
	State      map[string]interface{}
}

// Position represents a grid position.
type Position struct {
	X, Y, Z int // Z = era/temporal layer
}

// NewEntity creates a new entity.
func NewEntity(id, entityType string, pos Position) *Entity {
	return &Entity{
		ID:         id,
		Type:       entityType,
		Position:   pos,
		Attributes: NewAttributes(),
		State:      make(map[string]interface{}),
	}
}

// Move moves the entity to a new position.
func (e *Entity) Move(pos Position) {
	e.Position = pos
}

// DistanceTo returns the Manhattan distance to another entity.
func (e *Entity) DistanceTo(other *Entity) int {
	dx := e.Position.X - other.Position.X
	if dx < 0 {
		dx = -dx
	}
	dy := e.Position.Y - other.Position.Y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}
