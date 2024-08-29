package main

func FinalFloor(instructions string) int {
	floor := 0

	for _, char := range instructions {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}

	return floor
}
