package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Parse the input file and build the graph
	graph, err := ParseInput(file)
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	// Find and print the shortest route distance
	shortestDistance := FindShortestRoute(graph)
	fmt.Println("Shortest route distance:", shortestDistance)
}
