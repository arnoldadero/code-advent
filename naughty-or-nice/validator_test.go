package main

import "testing"

func TestIsNiceString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},  // nice: repeating pair (qj), repeat with one letter between (zxz)
		{"xxyxx", true},             // nice: repeating pair (xx), repeat with one letter between (yxy)
		{"uurcxstgmygtbstg", false}, // naughty: repeating pair (tg), but no repeat with one between
		{"ieodomkazucvgmuy", false}, // naughty: repeat with one between (odo), but no repeating pair
	}

	for _, test := range tests {
		result := IsNiceString(test.input)
		if result != test.expected {
			t.Errorf("IsNiceString(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestContainsRepeatingPair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyxy", true},       // repeating pair (xy)
		{"aabcdefgaa", true}, // repeating pair (aa)
		{"aaa", false},       // no non-overlapping pair
	}

	for _, test := range tests {
		if containsRepeatingPair(test.input) != test.expected {
			t.Errorf("containsRepeatingPair(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}

func TestContainsRepeatWithOneBetween(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyx", true},        // repeat with one letter between (xyx)
		{"abcdefeghi", true}, // repeat with one letter between (efe)
		{"aaa", true},        // repeat with one letter between (aaa)
		{"abc", false},       // no such repeating pattern
	}

	for _, test := range tests {
		if containsRepeatWithOneBetween(test.input) != test.expected {
			t.Errorf("containsRepeatWithOneBetween(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}
