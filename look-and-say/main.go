package main

import (
	"fmt"
)

func main() {
	// Initial input sequence
	input := "1321131112"

	// Number of iterations
	iterations := 40

	// Generate the look-and-say sequence for the given number of iterations
	result := GenerateSequence(input, iterations)

	// Print the length of the final result
	fmt.Printf("Length after %d iterations: %d\n", iterations, len(result))
}
