package main

import (
	"bufio"
	"os"
	"strings"
)

// ReadDimensions reads the dimensions of presents from a file and returns them as a slice of strings.
func ReadDimensions(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var dimensions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dimensions = append(dimensions, strings.TrimSpace(scanner.Text()))
	}
	return dimensions, scanner.Err()
}
