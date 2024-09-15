package main

import (
	"testing"
)

func TestNextSequence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, test := range tests {
		result := nextSequence(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}

func TestGenerateSequence(t *testing.T) {
	input := "1"
	iterations := 5
	expected := "312211"
	result := GenerateSequence(input, iterations)

	if result != expected {
		t.Errorf("For input %s after %d iterations, expected %s, but got %s", input, iterations, expected, result)
	}
}
