// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package game provides gameplay systems for Viva Las Mesh Layer 2.
// This is Layer 2 (Game Mechanics) - Gameplay systems (attrs, items, NPCs, missions).
package game

import (
	"time"
)

// RFGBroadcastChannel represents the Radio Free Graceland airwave broadcast channel.
type RFGBroadcastChannel struct {
	// Frequencies on which the channel broadcasts (in Hz)
	Frequencies []float64
	// isActive indicates whether the channel is currently broadcasting
	isActive bool
	// BroadcastTicker is used to simulate regular broadcasts
	BroadcastTicker *time.Ticker
	// DoneChan signals the ticker to stop
	DoneChan chan struct{}
	// LastBroadcastTime tracks when the last broadcast occurred
	LastBroadcastTime time.Time
}

// NewRFGBroadcastChannel creates a new Radio Free Graceland broadcast channel
// with the specified frequencies.
func NewRFGBroadcastChannel(freqs []float64) *RFGBroadcastChannel {
	rfc := &RFGBroadcastChannel{
		Frequencies:     freqs,
		isActive:        false,
		BroadcastTicker: time.NewTicker(10 * time.Second), // Broadcast every 10 seconds (example)
		DoneChan:        make(chan struct{}),
	}
	// Start the ticker to simulate broadcasts
	go rfc.startBroadcasting()
	return rfc
}

// startBroadcasting runs in a goroutine to simulate regular broadcasts.
func (rfc *RFGBroadcastChannel) startBroadcasting() {
	for {
		select {
		case <-rfc.BroadcastTicker.C:
			rfc.LastBroadcastTime = time.Now()
			// In a real implementation, this would broadcast audio/data on the frequencies
			// For now, we just log or do nothing; the effect is that the channel is active
		case <-rfc.DoneChan:
			rfc.BroadcastTicker.Stop()
			return
		}
	}
}

// Start begins broadcasting on the channel.
func (rfc *RFGBroadcastChannel) Start() {
	rfc.isActive = true
	// Ensure the ticker is running (it should already be started in NewRFGBroadcastChannel)
}

// Stop stops broadcasting.
func (rfc *RFGBroadcastChannel) Stop() {
	rfc.isActive = false
	rfc.DoneChan <- struct{}{}
}

// IsActive returns true if the channel is currently broadcasting.
func (rfc *RFGBroadcastChannel) IsActive() bool {
	return rfc.isActive
}

// GetFrequencies returns the list of frequencies on which the channel broadcasts.
func (rfc *RFGBroadcastChannel) GetFrequencies() []float64 {
	return rfc.Frequencies
}

// GetLastBroadcastTime returns the time of the last broadcast.
func (rfc *RFGBroadcastChannel) GetLastBroadcastTime() time.Time {
	return rfc.LastBroadcastTime
}
