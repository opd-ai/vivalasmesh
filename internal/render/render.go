// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package render provides sub-pixel rendering for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Sub-pixel rendering, Z-buffer, UI.
//
// Implements:
//   - Half-block (▀ U+2580) canvas (120×80 effective)
//   - 5-layer Z-buffer (Z-0 floor → Z-4 HUD)
//   - TrueColor ANSI escape sequence rendering
//   - Sub-pixel color blending (fg=top, bg=bottom)
//   - Particle/FX layer (Z-3)
//
// This package MUST NOT import any Layer 1 (transport/crypto/sync/daemon) packages.
package render

import (
	"errors"
	"image/color"
)

// Common errors.
var (
	ErrInvalidCanvasSize = errors.New("invalid canvas size")
	ErrInvalidLayer      = errors.New("invalid Z-buffer layer")
	ErrInvalidColor      = errors.New("invalid color")
)

// ZLayer represents a Z-buffer layer depth (0-4).
type ZLayer int

const (
	Z0Floor     ZLayer = iota // Floor/background
	Z1Terrain                 // Terrain/structures
	Z2Entities                // Entities/characters
	Z3Particles               // Particles/effects
	Z4HUD                     // HUD/spectrum analyzers
)

// String returns the layer name.
func (z ZLayer) String() string {
	switch z {
	case Z0Floor:
		return "Z-0:Floor"
	case Z1Terrain:
		return "Z-1:Terrain"
	case Z2Entities:
		return "Z-2:Entities"
	case Z3Particles:
		return "Z-3:Particles"
	case Z4HUD:
		return "Z-4:HUD"
	default:
		return "Unknown"
	}
}

// SubPixel represents a half-block sub-pixel cell.
// Top sub-pixel = foreground color, Bottom sub-pixel = background color.
type SubPixel struct {
	Top    color.RGBA // Upper half-block (▀ foreground)
	Bottom color.RGBA // Lower half-block (▀ background)
}

// Canvas is the sub-pixel rendering canvas (120×80 effective = 120×40 cells).
type Canvas struct {
	Width  int // Character cells wide (120)
	Height int // Character cells high (40)
	Layers [5][][]SubPixel
}

// NewCanvas creates a new sub-pixel canvas.
func NewCanvas(width, height int) (*Canvas, error) {
	if width <= 0 || height <= 0 {
		return nil, ErrInvalidCanvasSize
	}
	if width > 240 || height > 80 {
		return nil, ErrInvalidCanvasSize
	}

	c := &Canvas{
		Width:  width,
		Height: height,
	}

	for i := range c.Layers {
		c.Layers[i] = make([][]SubPixel, height)
		for y := range c.Layers[i] {
			c.Layers[i][y] = make([]SubPixel, width)
		}
	}

	return c, nil
}

// Set sets a sub-pixel at the given layer and position.
func (c *Canvas) Set(layer ZLayer, x, y int, sp SubPixel) error {
	if layer < Z0Floor || layer > Z4HUD {
		return ErrInvalidLayer
	}
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return ErrInvalidCanvasSize
	}
	c.Layers[layer][y][x] = sp
	return nil
}

// Get returns the sub-pixel at the given layer and position.
func (c *Canvas) Get(layer ZLayer, x, y int) (SubPixel, error) {
	if layer < Z0Floor || layer > Z4HUD {
		return SubPixel{}, ErrInvalidLayer
	}
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return SubPixel{}, ErrInvalidCanvasSize
	}
	return c.Layers[layer][y][x], nil
}

// Clear clears a specific layer.
func (c *Canvas) Clear(layer ZLayer) error {
	if layer < Z0Floor || layer > Z4HUD {
		return ErrInvalidLayer
	}
	for y := range c.Layers[layer] {
		for x := range c.Layers[layer][y] {
			c.Layers[layer][y][x] = SubPixel{}
		}
	}
	return nil
}

// ClearAll clears all layers.
func (c *Canvas) ClearAll() {
	for i := range c.Layers {
		c.Clear(ZLayer(i))
	}
}

// Render composites all layers and returns ANSI escape sequences.
// Returns a string ready to write to terminal.
func (c *Canvas) Render() string {
	// This is a simplified implementation
	// Real implementation would composite Z-layers and emit ANSI SGR sequences
	var result string
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			// Find topmost non-empty layer
			for layer := Z4HUD; layer >= Z0Floor; layer-- {
				sp := c.Layers[layer][y][x]
				if sp.Top != (color.RGBA{}) || sp.Bottom != (color.RGBA{}) {
					result += sp.ANSISequence()
					break
				}
			}
		}
		result += "\n"
	}
	return result
}

// ANSISequence returns the ANSI escape sequence for this sub-pixel.
func (sp SubPixel) ANSISequence() string {
	// Half-block rendering: fg=top, bg=bottom
	// \x1b[38;2;R;G;Bm\x1b[48;2;R;G;Bm▀
	if sp.Top == (color.RGBA{}) && sp.Bottom == (color.RGBA{}) {
		return " "
	}
	return "\x1b[38;2;" + rgb(sp.Top) + "m\x1b[48;2;" + rgb(sp.Bottom) + "m▀\x1b[0m"
}

func rgb(c color.RGBA) string {
	return itoa(int(c.R)) + ";" + itoa(int(c.G)) + ";" + itoa(int(c.B))
}

func itoa(i int) string {
	// Simple integer to string conversion
	if i == 0 {
		return "0"
	}
	var result string
	neg := false
	if i < 0 {
		neg = true
		i = -i
	}
	for i > 0 {
		result = string(rune('0'+i%10)) + result
		i /= 10
	}
	if neg {
		result = "-" + result
	}
	return result
}

// Color helpers
func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: a}
}
