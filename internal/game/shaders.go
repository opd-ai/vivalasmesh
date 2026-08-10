// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"image/color"
)

// ApplyParanoiaShader applies a color distortion shader based on the paranoia state.
// It takes an input color and returns a modified color that reflects the visual
// effect of the given paranoia state on the terminal canvas.
// In a real implementation, this would involve more complex color transformations
// to simulate the desired visual effects (e.g., hue shifts, saturation changes,
// brightness alterations, etc.).
func ApplyParanoiaShader(state ParanoiaState, src color.RGBA) color.RGBA {
	// We'll work with floating point values for better color manipulation
	r := float64(src.R) / 255.0
	g := float64(src.G) / 255.0
	b := float64(src.B) / 255.0

	switch state {
	case ParanoiaStateSmoothSynergy:
		// Smooth Synergy (1–5): slight warm tint, increase red and green slightly
		r = clamp01(r * 1.05)
		g = clamp01(g * 1.03)
		// blue unchanged
	case ParanoiaStateGreeningOut:
		// Greening Out (6–12): strong green tint, increase green, decrease red and blue
		r = clamp01(r * 0.8)
		g = clamp01(g * 1.2)
		b = clamp01(b * 0.8)
	case ParanoiaStateBatCountry:
		// Bat Country (13–20): dark, desaturated, with a yellow-brown tint
		// Reduce overall brightness, shift towards yellow
		r = clamp01(r*0.9 + 0.1)
		g = clamp01(g*0.9 + 0.1)
		b = clamp01(b * 0.7) // reduce blue more
	case ParanoiaStateExistentialDread:
		// Existential Dread (21+): high contrast, inverted colors, severe distortion
		// Invert colors and increase contrast
		r = clamp01(1.0 - r)
		g = clamp01(1.0 - g)
		b = clamp01(1.0 - b)
		// Increase contrast by pushing values away from 0.5
		r = clamp01((r-0.5)*1.5 + 0.5)
		g = clamp01((g-0.5)*1.5 + 0.5)
		b = clamp01((b-0.5)*1.5 + 0.5)
	default:
		// Normal state: no alteration
		return src
	}

	// Convert back to 0-255 range
	return color.RGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: src.A,
	}
}

// clamp01 ensures a float64 stays between 0 and 1.
func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// ParanoiaState represents the different paranoia states derived from the ParanoiaIndex.
type ParanoiaState string

const (
	ParanoiaStateNormal           ParanoiaState = "Normal"
	ParanoiaStateSmoothSynergy    ParanoiaState = "Smooth Synergy"    // 1–5
	ParanoiaStateGreeningOut      ParanoiaState = "Greening Out"      // 6–12
	ParanoiaStateBatCountry       ParanoiaState = "Bat Country"       // 13–20
	ParanoiaStateExistentialDread ParanoiaState = "Existential Dread" // 21+
)

// ParanoiaStateFromIndex converts a ParanoiaIndex to a ParanoiaState.
func ParanoiaStateFromIndex(index ParanoiaIndex) ParanoiaState {
	switch {
	case index >= 1 && index <= 5:
		return ParanoiaStateSmoothSynergy
	case index >= 6 && index <= 12:
		return ParanoiaStateGreeningOut
	case index >= 13 && index <= 20:
		return ParanoiaStateBatCountry
	case index >= 21:
		return ParanoiaStateExistentialDread
	default:
		return ParanoiaStateNormal
	}
}
