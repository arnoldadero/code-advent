package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Grid represents the 1000x1000 grid of lights with brightness levels
type Grid struct {
	brightness [][]int
}

// NewGrid creates a new 1000x1000 grid with all lights having brightness 0
func NewGrid() *Grid {
	brightness := make([][]int, 1000)
	for i := range brightness {
		brightness[i] = make([]int, 1000)
	}
	return &Grid{brightness}
}

// IncreaseBrightness increases the brightness of lights in the given range by the specified amount
func (g *Grid) IncreaseBrightness(x1, y1, x2, y2, amount int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			g.brightness[i][j] += amount
		}
	}
}

// DecreaseBrightness decreases the brightness of lights in the given range by the specified amount, but not below zero
func (g *Grid) DecreaseBrightness(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if g.brightness[i][j] > 0 {
				g.brightness[i][j]--
			}
		}
	}
}

// TotalBrightness calculates the total brightness of all the lights
func (g *Grid) TotalBrightness() int {
	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			total += g.brightness[i][j]
		}
	}
	return total
}

// ProcessInstructions processes the list of instructions on the grid
func ProcessInstructions(g *Grid, instructions []string) {
	for _, instruction := range instructions {
		parts := strings.Fields(instruction)

		var x1, y1, x2, y2 int

		if parts[0] == "toggle" {
			// toggle x1,y1 through x2,y2
			coords1 := strings.Split(parts[1], ",")
			coords2 := strings.Split(parts[3], ",")
			x1, _ = strconv.Atoi(coords1[0])
			y1, _ = strconv.Atoi(coords1[1])
			x2, _ = strconv.Atoi(coords2[0])
			y2, _ = strconv.Atoi(coords2[1])
			g.IncreaseBrightness(x1, y1, x2, y2, 2)
		} else {
			// turn on or off x1,y1 through x2,y2
			coords1 := strings.Split(parts[2], ",")
			coords2 := strings.Split(parts[4], ",")
			x1, _ = strconv.Atoi(coords1[0])
			y1, _ = strconv.Atoi(coords1[1])
			x2, _ = strconv.Atoi(coords2[0])
			y2, _ = strconv.Atoi(coords2[1])

			if parts[1] == "on" {
				g.IncreaseBrightness(x1, y1, x2, y2, 1)
			} else {
				g.DecreaseBrightness(x1, y1, x2, y2)
			}
		}
	}
}

func main() {
	g := NewGrid()

	// Reading the instructions from a file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	ProcessInstructions(g, instructions)
	fmt.Println("Total brightness:", g.TotalBrightness())
}
