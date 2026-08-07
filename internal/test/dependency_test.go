package internal_test

import (
	"os"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ipfs/go-datastore"
	"golang.org/x/crypto"
)

// TestDependenciesMatchSpecification validates that go.mod declares all dependencies required by
// the specification (AUDIT-001).
func TestDependenciesMatchSpecification(t *testing.T) {
	// Expected dependencies from ROADMAP.md Section 1.1 & spec compliance
	required := []string{
		"github.com/charmbracelet/bubbletea v1.3.10",
		"github.com/charmbracelet/lipgloss v1.1.0",
		"github.com/charmbracelet/bubbles v1.0.0",
		"github.com/ipfs/go-datastore v0.9.0",
		"golang.org/x/crypto v0.29.0",
	}

	// Read go.mod file
	content, err := os.ReadFile("../../go.mod")
	if err != nil {
		t.Fatalf("failed to read go.mod: %v", err)
	}

	for _, dep := range required {
		if !strings.Contains(string(content), dep) {
			t.Errorf("Missing or version mismatch for required dependency: %s", dep)
		}
	}

	// Implicit import check that dependencies resolve correctly
	_ = bubbletea.NewProgram(nil)
	_ = lipgloss.NewStyle()
	_ = datastore.NewMapDatastore
	_ = crypto.Salsa20
}
