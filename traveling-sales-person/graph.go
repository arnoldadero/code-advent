package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Graph represents the distances between cities
type Graph map[string]map[string]int

// ParseInput reads the input file and creates a graph of distances between cities
func ParseInput(r io.Reader) (Graph, error) {
	graph := make(Graph)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 5 {
			return nil, errors.New("invalid line format")
		}

		city1, city2 := parts[0], parts[2]
		distance, err := strconv.Atoi(parts[4])
		if err != nil {
			return nil, fmt.Errorf("invalid distance format: %v", err)
		}

		// Add to graph in both directions
		if graph[city1] == nil {
			graph[city1] = make(map[string]int)
		}
		graph[city1][city2] = distance

		if graph[city2] == nil {
			graph[city2] = make(map[string]int)
		}
		graph[city2][city1] = distance
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %v", err)
	}

	return graph, nil
}
