package main

import "testing"

func TestCountUniqueHouses(t *testing.T) {
	tests := []struct {
		moves    string
		expected int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
		{"^^vv<<>>", 5},
	}

	for _, test := range tests {
		result := CountUniqueHouses(test.moves)
		if result != test.expected {
			t.Errorf("For moves '%s', expected %d but got %d", test.moves, test.expected, result)
		}
	}
}
