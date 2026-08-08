// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package types provides shared type definitions for Viva Las Mesh.
package types

// Shared types

// Color represents an RGBA color.
type Color struct {
	R, G, B, A uint8
}

// Rect represents a rectangle.
type Rect struct {
	X, Y, W, H int
}

// Contains checks if a point is in the rectangle.
func (r Rect) Contains(x, y int) bool {
	return x >= r.X && x < r.X+r.W && y >= r.Y && y < r.Y+r.H
}

// Vector2 represents a 2D vector.
type Vector2 struct {
	X, Y float64
}

// Add returns the sum of two vectors.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// Sub returns the difference of two vectors.
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

// Mul returns the vector scaled by a scalar.
func (v Vector2) Mul(scalar float64) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}
