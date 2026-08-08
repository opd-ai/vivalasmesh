// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package render provides rendering sub-systems for Viva Las Mesh.
// Implements sub-pixel terminal rendering using upper half-block characters (���▀)
// and a 5-layer Z-buffer compositing pipeline.
package render

import (
	"fmt"
	"image/color"
	"sync"
)

// Color is an alias for color.RGBA for convenience.
type Color = color.RGBA

// Cell represents a single terminal cell for sub-pixel rendering.
// The upper half of the cell is displayed in the foreground color,
// the lower half in the background color, using the '���▀' (U+2580) character.
type Cell struct {
	Fg Color
	Bg Color
}

// Layer is a 2D grid of color pointers, where nil represents transparency.
type Layer [][]*Color

// Renderer holds the state for sub-pixel rendering.
// It manages a 5-layer Z-buffer and provides methods to draw, composite, and render.
type Renderer struct {
	mu      sync.RWMutex
	width   int // width in terminal cells (columns)
	height  int // height in terminal cells (rows)
	layers  [5]Layer
}

// NewRenderer creates a new Renderer with the given dimensions.
// width and height are in terminal cells. Each cell represents two sub-pixel rows
// (upper and lower half), so the effective sub-pixel resolution is width x (height*2).
func NewRenderer(width, height int) *Renderer {
	r := &Renderer{
		width:  width,
		height: height,
	}
	for i := range r.layers {
		r.layers[i] = make(Layer, height)
		for y := range r.layers[i] {
			r.layers[i][y] = make([]*Color, width)
		}
	}
	return r
}

// SetPixel sets the color of a single cell in the specified layer.
// Layer index 0 is the bottom layer (Z-0), 4 is the top layer (Z-4).
// If color is nil, the cell becomes transparent in that layer.
func (r *Renderer) SetPixel(layer int, x, y int, c *Color) {
	if layer < 0 || layer > 4 {
		return
	}
	if x < 0 || x >= r.width || y < 0 || y >= r.height {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.layers[layer][y][x] = c
}

// ClearLayer sets all cells in the specified layer to transparent.
func (r *Renderer) ClearLayer(layer int) {
	if layer < 0 || layer > 4 {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			r.layers[layer][y][x] = nil
		}
	}
}

// Composite combines all layers into a single buffer of Cells using the Z-buffer algorithm.
// For each cell, the topmost non-transparent layer's color is used for both foreground and background.
// Returns a 2D slice of Cells representing the final image to be rendered.
func (r *Renderer) Composite() [][]Cell {
	r.mu.RLock()
	defer r.mu.RUnlock()
	buf := make([][]Cell, r.height)
	for y := 0; y < r.height; buf[y] = make([]Cell, r.width), y++ {
		for x := 0; x < r.width; x++ {
			var c *Color
			// Check layers from top (4) to bottom (0) to find the topmost non-transparent color.
			for i := 4; i >= 0; i-- {
				if col := r.layers[i][y][x]; col != nil {
					c = col
					break
				}
			}
			if c == nil {
				// Transparent: set both colors to transparent (alpha 0).
				buf[y][x] = Cell{Fg: Color{A: 0}, Bg: Color{A: 0}}
			} else {
				// Opaque: use the same color for foreground and background.
				buf[y][x] = Cell{Fg: *c, Bg: *c}
			}
		}
	}
	return buf
}

// Render returns an ANSI escape sequence string representing the rendered buffer.
// Each terminal cell is rendered as the '���▀' (U+2580) character with the
// appropriate foreground and background colors set via 24-bit ANSI escape codes.
// The string includes a carriage return and newline at the end of each row,
// and reset codes after each row to prevent color bleeding.
func (r *Renderer) Render() string {
	buf := r.Composite()
	var out string
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			cell := buf[y][x]
			// Only output color escapes if the color is opaque (alpha == 255).
			// We'll output the escape sequences regardless of alpha for simplicity.
			fgEsc := ""
			bgEsc := ""
			if cell.Fg.A != 0 {
				fgEsc = fmt.Sprintf("\x1b[38;2;%d;%d;%dm", cell.Fg.R, cell.Fg.G, cell.Fg.B)
			}
			if cell.Bg.A != 0 {
				bgEsc = fmt.Sprintf("\x1b[48;2;%d;%d;%dm", cell.Bg.R, cell.Bg.G, cell.Bg.B)
			}
			out += fgEsc + bgEsc + "���▀"
		}
		out += "\x1b[0m\r\n"
	}
	return out
}

// Resize updates the renderer's dimensions and reallocates the layers.
// width and height are in terminal cells.
// Resize updates the renderer's dimensions and reallocates the layers.
// width and height are in terminal cells.
func (r *Renderer) Resize(width, height int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.width = width
	r.height = height
	for i := range r.layers {
		r.layers[i] = make(Layer, height)
		for y := range r.layers[i] {
			r.layers[i][y] = make([]*Color, width)
		}
	}
}
