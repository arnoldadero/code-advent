package main

import (
	"strconv"
	"strings"
)

// GenerateSequence generates the look-and-say sequence for the given input
// for the specified number of iterations.
func GenerateSequence(input string, iterations int) string {
	current := input

	for i := 0; i < iterations; i++ {
		current = nextSequence(current)
	}

	return current
}

// nextSequence takes the current sequence and returns the next sequence
// in the look-and-say process.
func nextSequence(seq string) string {
	var builder strings.Builder
	count := 1

	for i := 1; i < len(seq); i++ {
		if seq[i] == seq[i-1] {
			count++
		} else {
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(seq[i-1])
			count = 1
		}
	}
	// Append the last group
	builder.WriteString(strconv.Itoa(count))
	builder.WriteByte(seq[len(seq)-1])

	return builder.String()
}
