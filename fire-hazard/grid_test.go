package main

import (
	"testing"
)

func TestGridBrightness(t *testing.T) {
	g := NewGrid()

	// Test turning on a single light
	g.IncreaseBrightness(0, 0, 0, 0, 1)
	if g.TotalBrightness() != 1 {
		t.Errorf("Expected total brightness 1, got %d", g.TotalBrightness())
	}

	// Test toggling the entire grid
	g.IncreaseBrightness(0, 0, 999, 999, 2)
	if g.TotalBrightness() != 2000001 {
		t.Errorf("Expected total brightness 2000001, got %d", g.TotalBrightness())
	}

	// Test turning off part of the grid
	g.DecreaseBrightness(499, 499, 500, 500)
	if g.TotalBrightness() != 1999997 {
		t.Errorf("Expected total brightness 1999999, got %d", g.TotalBrightness())
	}
}
