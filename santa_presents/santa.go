package main

// Point represents a coordinate on the grid
type Point struct {
	X int
	Y int
}

// CountUniqueHouses counts the number of unique houses that receive at least one present
func CountUniqueHouses(moves string) int {
	visitedHouses := make(map[Point]bool)
	currentPosition := Point{X: 0, Y: 0}
	visitedHouses[currentPosition] = true

	for _, move := range moves {
		switch move {
		case '^':
			currentPosition.Y++
		case 'v':
			currentPosition.Y--
		case '>':
			currentPosition.X++
		case '<':
			currentPosition.X--
		}
		visitedHouses[currentPosition] = true
	}

	return len(visitedHouses)
}

// CountUniqueHousesWithRoboSanta counts the number of unique houses that receive at least one present when both Santa and Robo-Santa are delivering presents
func CountUniqueHousesWithRoboSanta(moves string) int {
	visitedHouses := make(map[Point]bool)
	santaPosition := Point{X: 0, Y: 0}
	roboSantaPosition := Point{X: 0, Y: 0}
	visitedHouses[santaPosition] = true

	for i, move := range moves {
		var currentPosition *Point
		if i%2 == 0 {
			currentPosition = &santaPosition
		} else {
			currentPosition = &roboSantaPosition
		}

		switch move {
		case '^':
			currentPosition.Y++
		case 'v':
			currentPosition.Y--
		case '>':
			currentPosition.X++
		case '<':
			currentPosition.X--
		}
		visitedHouses[*currentPosition] = true
	}

	return len(visitedHouses)
}