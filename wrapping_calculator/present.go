package main

import (
	"sort"
	"strconv"
	"strings"
)

// Present represents a box with dimensions.
type Present struct {
	Length int
	Width  int
	Height int
}

// NewPresent creates a new Present instance from a string of dimensions "LxWxH".
func NewPresent(dimensions string) Present {
	dim := strings.Split(dimensions, "x")
	length, _ := strconv.Atoi(dim[0])
	width, _ := strconv.Atoi(dim[1])
	height, _ := strconv.Atoi(dim[2])
	return Present{Length: length, Width: width, Height: height}
}

// WrappingPaperRequired calculates the wrapping paper required for a single present.
func (p Present) WrappingPaperRequired() int {
	surfaceArea := 2*p.Length*p.Width + 2*p.Width*p.Height + 2*p.Height*p.Length
	sides := []int{p.Length * p.Width, p.Width * p.Height, p.Height * p.Length}
	sort.Ints(sides)
	return surfaceArea + sides[0] // Adding the smallest side as extra paper
}
