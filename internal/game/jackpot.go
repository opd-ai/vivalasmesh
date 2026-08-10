// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"sync"
	"time"
)

// GlobalJackpot represents the Global 777 CRDT Progressive Jackpot.
// It aggregates casino losses and heist failures into a cross-network pool.
// In a real implementation, this would be a CRDT that converges across network nodes.
// For simplicity, we simulate it with a mutex-protected value.
type GlobalJackpot struct {
	// currentValue is the current jackpot amount
	currentValue int64
	// mutex protects access to currentValue
	mutex sync.RWMutex
	// lastUpdateTime tracks when the jackpot was last updated
	lastUpdateTime time.Time
	// winThreshold is the amount at which the jackpot can be won
	// In a real implementation, this might be dynamic or based on a random roll
	winThreshold int64
	// won indicates whether the jackpot has been won (and needs to reset)
	won bool
}

// NewGlobalJackpot creates a new Global Jackpot with an initial seed value.
func NewGlobalJackpot(seed int64) *GlobalJackpot {
	return &GlobalJackpot{
		currentValue:   seed,
		winThreshold:   777777, // Example threshold
		won:            false,
		lastUpdateTime: time.Now(),
	}
}

// AddLoss adds an amount to the jackpot from a casino loss or heist failure.
// This is thread-safe.
func (gj *GlobalJackpot) AddLoss(amount int64) {
	gj.mutex.Lock()
	defer gj.mutex.Unlock()
	gj.currentValue += amount
	gj.lastUpdateTime = time.Now()
	// Check if the jackpot has reached or exceeded the win threshold
	if gj.currentValue >= gj.winThreshold && !gj.won {
		// In a real implementation, this would trigger a win event via CRDT
		// For now, we just mark it as won; the actual win would be determined by gameplay
		gj.won = true
	}
}

// GetValue returns the current jackpot value.
func (gj *GlobalJackpot) GetValue() int64 {
	gj.mutex.RLock()
	defer gj.mutex.RUnlock()
	return gj.currentValue
}

// IsWon returns true if the jackpot has reached the win threshold and is ready to be won.
func (gj *GlobalJackpot) IsWon() bool {
	gj.mutex.RLock()
	defer gj.mutex.RUnlock()
	return gj.won
}

// WinJackpot resets the jackpot to a seed value after being won.
// This should be called when a player successfully wins the jackpot.
func (gj *GlobalJackpot) WinJackpot(seed int64) {
	gj.mutex.Lock()
	defer gj.mutex.Unlock()
	gj.currentValue = seed
	gj.won = false
	gj.lastUpdateTime = time.Now()
}

// GetLastUpdateTime returns the time of the last update to the jackpot.
func (gj *GlobalJackpot) GetLastUpdateTime() time.Time {
	gj.mutex.RLock()
	defer gj.mutex.RUnlock()
	return gj.lastUpdateTime
}

// SetWinThreshold sets the amount at which the jackpot can be won.
func (gj *GlobalJackpot) SetWinThreshold(threshold int64) {
	gj.mutex.Lock()
	defer gj.mutex.Unlock()
	gj.winThreshold = threshold
}

// GetWinThreshold returns the current win threshold.
func (gj *GlobalJackpot) GetWinThreshold() int64 {
	gj.mutex.RLock()
	defer gj.mutex.RUnlock()
	return gj.winThreshold
}
