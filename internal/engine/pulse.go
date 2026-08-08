// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

// Package engine provides the game engine core for Viva Las Mesh Layer 2.
package engine

import (
	"context"
	"sync"
	"time"
)

// PulseEngine manages the 250ms micro-tick and 1000ms macro-pulse.
type PulseEngine struct {
	ctx          context.Context
	cancel       context.CancelFunc
	microTick    time.Duration
	macroPulse   time.Duration
	microTicks   int
	actionQueue  *ActionQueue
	running      bool
	runningMu    sync.Mutex
	wg           sync.WaitGroup
	onMicroTick  func(int)
	onMacroPulse func()
}

// NewPulseEngine creates a new pulse engine.
func NewPulseEngine(microTick, macroPulse time.Duration, queueSize int) *PulseEngine {
	ctx, cancel := context.WithCancel(context.Background())
	return &PulseEngine{
		ctx:         ctx,
		cancel:      cancel,
		microTick:   microTick,
		macroPulse:  macroPulse,
		actionQueue: NewActionQueue(queueSize),
	}
}

// Start starts the pulse engine.
func (p *PulseEngine) Start() error {
	p.runningMu.Lock()
	defer p.runningMu.Unlock()

	if p.running {
		return nil
	}

	p.running = true
	p.wg.Add(1)
	go p.run()
	return nil
}

// Stop stops the pulse engine.
func (p *PulseEngine) Stop() error {
	p.runningMu.Lock()
	defer p.runningMu.Unlock()

	if !p.running {
		return ErrPulseNotRunning
	}

	p.cancel()
	p.wg.Wait()
	p.running = false
	return nil
}

// StageAction stages an action for the next pulse.
func (p *PulseEngine) StageAction(action *Action) error {
	return p.actionQueue.Stage(action)
}

// run is the main pulse loop.
func (p *PulseEngine) run() {
	defer p.wg.Done()

	microTicker := time.NewTicker(p.microTick)
	macroTicker := time.NewTicker(p.macroPulse)
	defer microTicker.Stop()
	defer macroTicker.Stop()

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-microTicker.C:
			p.microTicks++
			if p.onMicroTick != nil {
				p.onMicroTick(p.microTicks)
			}

			// Every 4 micro-ticks = 1 macro-pulse (1000ms)
			if p.microTicks%4 == 0 {
				resolved := p.actionQueue.Resolve()
				// Process resolved actions
				_ = resolved

				if p.onMacroPulse != nil {
					p.onMacroPulse()
				}
			}
		case <-macroTicker.C:
			// Macro pulse fired
			if p.onMacroPulse != nil {
				p.onMacroPulse()
			}
		}
	}
}

// SetMicroTickHandler sets the micro-tick callback.
func (p *PulseEngine) SetMicroTickHandler(fn func(int)) {
	p.onMicroTick = fn
}

// SetMacroPulseHandler sets the macro-pulse callback.
func (p *PulseEngine) SetMacroPulseHandler(fn func()) {
	p.onMacroPulse = fn
}

// MicroTicks returns the current micro-tick count.
func (p *PulseEngine) MicroTicks() int {
	return p.microTicks
}
