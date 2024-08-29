package main

import (
	"fmt"
	"log"
)

func main() {
	// Read the dimensions from the input file
	presents, err := ReadDimensions("presents.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Calculate the total wrapping paper required
	totalWrappingPaper := 0
	totalRibbon := 0
	for _, dimensions := range presents {
		p := NewPresent(dimensions)
		totalWrappingPaper += p.WrappingPaperRequired()
		totalRibbon += p.RibbonRequired()
	}

	fmt.Printf("Total wrapping paper required: %d square feet\n", totalWrappingPaper)
	fmt.Printf("Total ribbon required: %d feet\n", totalRibbon)
}
