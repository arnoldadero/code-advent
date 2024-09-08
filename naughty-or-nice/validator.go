package main

import (
	"strings"
)

// IsNiceString checks if a string meets the new "nice" criteria.
func IsNiceString(str string) bool {
	return containsRepeatingPair(str) && containsRepeatWithOneBetween(str)
}

// containsRepeatingPair checks if the string contains a pair of any two letters
// that appears at least twice without overlapping.
func containsRepeatingPair(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		pair := str[i : i+2]
		// Check if this pair appears again later in the string
		if strings.Contains(str[i+2:], pair) {
			return true
		}
	}
	return false
}

// containsRepeatWithOneBetween checks if the string contains at least one letter
// that repeats with exactly one letter between them (e.g., xyx).
func containsRepeatWithOneBetween(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}
