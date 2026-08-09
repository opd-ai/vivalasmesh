// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

package release

// TargetPlatform represents a cross-compilation target.
type TargetPlatform struct {
	GOOS   string
	GOARCH string
}

// DefaultTargets returns the default cross-compilation targets for production releases.
func DefaultTargets() []TargetPlatform {
	return []TargetPlatform{
		{GOOS: "linux", GOARCH: "amd64"},
		{GOOS: "linux", GOARCH: "arm64"},
		{GOOS: "darwin", GOARCH: "amd64"}, // Apple Silicon
		{GOOS: "darwin", GOARCH: "arm64"}, // Intel Mac
		{GOOS: "windows", GOARCH: "amd64"},
	}
}
