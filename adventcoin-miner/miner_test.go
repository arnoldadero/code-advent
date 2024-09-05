package main

import (
	"testing"
)

func TestMineAdventCoin(t *testing.T) {
	testCases := []struct {
		secretKey    string
		prefixLength int
		expected     int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}

	for _, tc := range testCases {
		result, err := MineAdventCoin(tc.secretKey, tc.prefixLength)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if result != tc.expected {
			t.Errorf("For secretKey %s, expected %d but got %d", tc.secretKey, tc.expected, result)
		}
	}
}
