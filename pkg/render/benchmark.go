// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

package render

import (
	"fmt"
	"time"
)

// BenchmarkRender measures the average time to render a frame over n iterations.
// It returns the average duration per frame and an error if any.
func BenchmarkRender(r *Renderer, n int) (time.Duration, error) {
	if n <= 0 {
		return 0, fmt.Errorf("iteration count must be positive")
	}
	var total time.Duration
	for i := 0; i < n; i++ {
		start := time.Now()
		_ = r.Render() // discard result
		total += time.Since(start)
	}
	return total / time.Duration(n), nil
}

// MustBenchmarkRender panics if the benchmark returns an error.
func MustBenchmarkRender(r *Renderer, n int) time.Duration {
	d, err := BenchmarkRender(r, n)
	if err != nil {
		panic(err)
	}
	return d
}
