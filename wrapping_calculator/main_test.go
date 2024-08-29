package main

import "testing"

func TestWrappingPaperRequired(t *testing.T) {
	tests := []struct {
		dimensions string
		expected   int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, tt := range tests {
		p := NewPresent(tt.dimensions)
		result := p.WrappingPaperRequired()
		if result != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, result)
		}
	}
}

func TestReadDimensions(t *testing.T) {
	_, err := ReadDimensions("presents.txt")
	if err != nil {
		t.Errorf("Error reading dimensions from file: %v", err)
	}
}
