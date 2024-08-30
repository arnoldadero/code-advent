package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// Read the input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Calculate the number of houses that receive at least one present
	moves := string(data)
	housesCount := CountUniqueHouses(moves)

	// Output the result
	fmt.Printf("Number of unique houses that receive at least one present: %d\n", housesCount)
}
