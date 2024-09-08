package main

import (
	"fmt"
	"log"
)

func main() {
	// File path to the input file
	filePath := "input.txt"

	// Reading the strings from the input file
	strings, err := ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	// Counting nice strings
	niceCount := 0
	for _, str := range strings {
		if IsNiceString(str) {
			niceCount++
		}
	}

	fmt.Printf("Number of nice strings: %d\n", niceCount)
}
