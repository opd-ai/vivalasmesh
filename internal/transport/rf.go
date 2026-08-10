// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package transport provides real-world multi-transport engine for Viva Las Mesh Layer 1.
// This is Layer 1 (P2P Infrastructure) - Real-world multi-transport engine.
// This package MUST NOT import any Layer 2 (game engine) packages.
package transport

import (
	"context"
	"math/rand"
	"time"
)

// Receiver represents a full-spectrum RF receiver capable of tuning across
// frequencies from VLF (3 kHz) through Ku-Band (18 GHz) and Optical/IR frequencies.
// This is a simulated architecture for the game; in a real implementation,
// this would interface with actual SDR hardware or signal processing chains.
type Receiver interface {
	// Tune sets the receiver center frequency in Hz.
	Tune(ctx context.Context, freqHz float64) error
	// GetSpectrum returns normalized amplitude values for frequency bins
	// around the tuned frequency. The number of bins and bin width are
	// implementation-specific.
	GetSpectrum() []float64
	// Start begins continuous reception and updates the internal spectrum buffer.
	Start(ctx context.Context) error
	// Stop stops reception.
	Stop()
}

// SimpleReceiver is a basic simulated RF receiver that generates
// synthetic spectrum data across the full spectrum.
type SimpleReceiver struct {
	// centerFreqHz is the current tuned frequency
	centerFreqHz float64
	// spectrum holds the normalized amplitude values for each frequency bin.
	// We'll simulate 120 bins covering a wide range (logarithmic scale).
	spectrum []float64
	// spectrumTicker is used to periodically update the spectrum.
	spectrumTicker *time.Ticker
	// doneChan is used to signal the receiver to stop.
	doneChan chan struct{}
}

// NewSimpleReceiver creates a new simple RF receiver with a default
// frequency range suitable for simulation.
func NewSimpleReceiver() *SimpleReceiver {
	// Initialize with 120 frequency bins (one per column in the TUI spectrum display)
	sr := &SimpleReceiver{
		spectrum:       make([]float64, 120),
		spectrumTicker: time.NewTicker(100 * time.Millisecond),
		doneChan:       make(chan struct{}),
	}
	// Initialize spectrum with zeros
	for i := range sr.spectrum {
		sr.spectrum[i] = 0
	}
	return sr
}

// Tune sets the receiver center frequency.
// In this simple implementation, we ignore the frequency for simulation,
// but in a real receiver, this would adjust the local oscillator and filters.
func (sr *SimpleReceiver) Tune(_ context.Context, freqHz float64) error {
	sr.centerFreqHz = freqHz
	return nil
}

// GetSpectrum returns the current spectrum buffer.
func (sr *SimpleReceiver) GetSpectrum() []float64 {
	return sr.spectrum
}

// Start begins continuous reception and updates the internal spectrum buffer.
// It simulates receiving RF signals across the spectrum with some peaks and noise.
func (sr *SimpleReceiver) Start(_ context.Context) error {
	go func() {
		for {
			select {
			case <-sr.spectrumTicker.C:
				sr.generateSpectrum()
			case <-sr.doneChan:
				sr.spectrumTicker.Stop()
				return
			}
		}
	}()
	return nil
}

// Stop stops reception.
func (sr *SimpleReceiver) Stop() {
	close(sr.doneChan)
}

// generateSpectrum populates the spectrum buffer with simulated data.
// In a real implementation, this would come from actual RF signal processing.
func (sr *SimpleReceiver) generateSpectrum() {
	// Simulate some RF signals: random noise with a few peaks across the spectrum
	for i := range sr.spectrum {
		// Base noise
		sr.spectrum[i] = rand.Float64() * 0.3
		// Add a few simulated signals at different frequencies
		// We'll map frequencies to bins logarithmically for wide range simulation
		// For simplicity, we'll just place peaks at fixed bins
		if i == 30 {
			sr.spectrum[i] = 0.8 + rand.Float64()*0.2
		}
		if i == 60 {
			sr.spectrum[i] = 0.6 + rand.Float64()*0.2
		}
		if i == 90 {
			sr.spectrum[i] = 0.9 + rand.Float64()*0.1
		}
		// Occasionally spike
		if rand.Float64() < 0.01 {
			sr.spectrum[i] = rand.Float64()*0.5 + 0.5
		}
	}
}
