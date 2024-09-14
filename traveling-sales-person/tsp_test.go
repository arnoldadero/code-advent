package main

import (
	"strings"
	"testing"
)

// TestParseInput verifies that the input is parsed correctly into a graph
func TestParseInput(t *testing.T) {
	input := `Faerun to Tristram = 65
Faerun to Tambi = 129
Faerun to Norrath = 144`

	graph, err := ParseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to parse input: %v", err)
	}

	// Check graph structure
	if graph["Faerun"]["Tristram"] != 65 || graph["Tristram"]["Faerun"] != 65 {
		t.Errorf("Faerun to Tristram distance incorrect: got %d", graph["Faerun"]["Tristram"])
	}
	if graph["Faerun"]["Tambi"] != 129 || graph["Tambi"]["Faerun"] != 129 {
		t.Errorf("Faerun to Tambi distance incorrect: got %d", graph["Faerun"]["Tambi"])
	}
	if graph["Faerun"]["Norrath"] != 144 || graph["Norrath"]["Faerun"] != 144 {
		t.Errorf("Faerun to Norrath distance incorrect: got %d", graph["Faerun"]["Norrath"])
	}
}

// TestPermutations checks if all permutations of the cities are generated
func TestPermutations(t *testing.T) {
	cities := []string{"A", "B", "C"}
	var permutations [][]string
	Permutations(cities, &permutations, 0)

	expectedCount := 6
	if len(permutations) != expectedCount {
		t.Errorf("Expected %d permutations, got %d", expectedCount, len(permutations))
	}
}

// TestCalculateRouteDistance ensures the distance calculation is correct
func TestCalculateRouteDistance(t *testing.T) {
	graph := Graph{
		"A": {"B": 5, "C": 10},
		"B": {"A": 5, "C": 7},
		"C": {"A": 10, "B": 7},
	}

	route := []string{"A", "B", "C"}
	expectedDistance := 12
	distance := CalculateRouteDistance(route, graph)
	if distance != expectedDistance {
		t.Errorf("Expected distance %d, got %d", expectedDistance, distance)
	}
}

// TestFindShortestRoute verifies the correct shortest route is found
func TestFindShortestRoute(t *testing.T) {
	graph := Graph{
		"A": {"B": 5, "C": 10},
		"B": {"A": 5, "C": 7},
		"C": {"A": 10, "B": 7},
	}

	expectedShortestDistance := 12
	shortestDistance := FindShortestRoute(graph)
	if shortestDistance != expectedShortestDistance {
		t.Errorf("Expected shortest distance %d, got %d", expectedShortestDistance, shortestDistance)
	}
}
