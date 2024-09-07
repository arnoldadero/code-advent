package main

import (
	"fmt"
	"log"
)

func main() {
	// File path to the input file
	filepath := "input.txt"

	// Reading the strings from the input file
	strings, err := ReadStrings(filepath)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Counting the nice strings
	niceCount := 0
	for _, str := range strings {
		if IsNiceString(str) {
			niceCount++
		}
	}
	fmt.Printf("Number of nice strings: %d\n", niceCount)
}
