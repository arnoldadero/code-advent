package main

import (
	"bufio"
	"os"
)

func ReadInstructions(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var instructions string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return instructions, nil
}
