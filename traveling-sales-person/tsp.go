package main

// CalculateRouteDistance computes the total distance for a specific route
func CalculateRouteDistance(route []string, graph Graph) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		city1, city2 := route[i], route[i+1]
		totalDistance += graph[city1][city2]
	}
	return totalDistance
}

// FindShortestRoute calculates the shortest route by evaluating all permutations
func FindShortestRoute(graph Graph) int {
	// Extract the list of cities
	var cities []string
	for city := range graph {
		cities = append(cities, city)
	}

	// Generate all possible routes (permutations)
	var permutations [][]string
	Permutations(cities, &permutations, 0)

	// Find the minimum route distance
	minDistance := -1
	for _, route := range permutations {
		distance := CalculateRouteDistance(route, graph)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

// FindLongestRoute calculates the longest route by evaluating all permutations
func FindLongestRoute(graph Graph) int {
	// Extract the list of cities
	var cities []string
	for city := range graph {
		cities = append(cities, city)
	}

	// Generate all possible routes (permutations)
	var permutations [][]string
	Permutations(cities, &permutations, 0)

	// Find the maximum route distance
	maxDistance := -1
	for _, route := range permutations {
		distance := CalculateRouteDistance(route, graph)
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	return maxDistance
}

// Permutations generates all permutations of the cities slice
func Permutations(cities []string, result *[][]string, start int) {
	if start == len(cities)-1 {
		// Make a copy of cities to avoid referencing issues
		tmp := make([]string, len(cities))
		copy(tmp, cities)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(cities); i++ {
		cities[start], cities[i] = cities[i], cities[start]
		Permutations(cities, result, start+1)
		cities[start], cities[i] = cities[i], cities[start] // backtrack
	}
}
