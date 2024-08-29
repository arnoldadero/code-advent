package main

func FinalFloor(instructions string) (int, int) {
	floor := 0
	basementPos := -1

	for i, char := range instructions {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		// Track When santa first enter the basement
		if basementPos == -1 && floor == -1 {
			basementPos = i + 1
		}
	}

	return floor, basementPos
}
