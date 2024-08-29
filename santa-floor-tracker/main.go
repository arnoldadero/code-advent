package main

import "fmt"

func main() {
	filename := "instructions.txt"
	instructions, err := ReadInstructions(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	final := FinalFloor(instructions)
	fmt.Printf("Santa ends up on floor: %d\n", final)
}