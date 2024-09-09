package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wires = make(map[string]string) // Stores the instructions for each wire
var cache = make(map[string]uint16) // Stores the evaluated signal for each wire

func main() {
	// Specify the input file
	inputFile := "input.txt"

	// Read instructions from the input file
	err := readInstructions(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Compute the signal for wire 'a'
	signalA := evaluate("a")
	fmt.Println("Initial signal on wire 'a':", signalA)

	// Override wire 'b' with the value of wire 'a'
	wires["b"] = fmt.Sprintf("%d", signalA)

	// Clear the cache and reset everything for re-evaluation
	cache = make(map[string]uint16)

	// Recompute the signal for wire 'a' with the new instruction for wire 'b'
	newSignalA := evaluate("a")
	fmt.Println("New signal on wire 'a' after overriding 'b':", newSignalA)
}

// readInstructions reads each line from the input file and stores the instructions in the wires map
func readInstructions(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseInstruction(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// parseInstruction parses a single line of input and stores it in the wires map
func parseInstruction(instruction string) {
	parts := strings.Split(instruction, " -> ")
	dest := parts[1]
	op := parts[0]
	wires[dest] = op
}

// evaluate recursively evaluates the signal on a wire
func evaluate(wire string) uint16 {
	// Check if the wire value is already cached
	if val, ok := cache[wire]; ok {
		return val
	}

	// Try to convert wire to a number (base case)
	if num, err := strconv.Atoi(wire); err == nil {
		return uint16(num)
	}

	// Get the operation for the wire
	op := wires[wire]
	parts := strings.Split(op, " ")

	var result uint16

	switch len(parts) {
	case 1:
		// Direct assignment
		result = evaluate(parts[0])

	case 2:
		// NOT operation
		result = ^evaluate(parts[1]) & 65535

	case 3:
		// Binary operations (AND, OR, LSHIFT, RSHIFT)
		left := evaluate(parts[0])
		right := evaluate(parts[2])

		switch parts[1] {
		case "AND":
			result = left & right
		case "OR":
			result = left | right
		case "LSHIFT":
			n, _ := strconv.Atoi(parts[2])
			result = left << n
		case "RSHIFT":
			n, _ := strconv.Atoi(parts[2])
			result = left >> n
		}
	}

	// Cache and return the result
	cache[wire] = result
	return result
}
