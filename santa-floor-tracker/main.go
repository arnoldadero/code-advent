package main

import "fmt"

func main() {
	filename := "instructions.txt"
	instructions, err := ReadInstructions(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	final, basementPos := FinalFloor(instructions)
	fmt.Printf("Santa ends up on floor: %d\n", final)
	if basementPos != -1 {
		fmt.Printf("Santa first enters the basement at position: %d\n", basementPos)
	} else {
		fmt.Println("Santa never enters the basement.")
	}
}
