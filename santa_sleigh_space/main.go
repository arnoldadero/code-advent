package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to calculate the difference between the original code length and the encoded length
func calculateEncodedDiff(filename string) (int, error) {
	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Initialize counters for total characters
	totalCodeChars := 0
	totalEncodedChars := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line and trim it (remove whitespace/newlines)
		line := strings.TrimSpace(scanner.Text())

		// Count characters in the string's code representation
		codeLength := len(line)
		totalCodeChars += codeLength

		// Get the encoded version of the string and count its length
		encodedLine := encodeString(line)
		encodedLength := len(encodedLine)
		totalEncodedChars += encodedLength
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Return the difference between total encoded characters and total code characters
	return totalEncodedChars - totalCodeChars, nil
}

// Helper function to encode the string by escaping backslashes and double quotes
func encodeString(s string) string {
	// Start with an opening quote
	var encodedString strings.Builder
	encodedString.WriteString("\"")

	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			encodedString.WriteString("\\\\") // Escape backslash
		} else if s[i] == '"' {
			encodedString.WriteString("\\\"") // Escape double quote
		} else {
			encodedString.WriteByte(s[i]) // Regular character
		}
	}

	// End with a closing quote
	encodedString.WriteString("\"")

	return encodedString.String()
}

func main() {
	// File containing the input strings
	filename := "input.txt"

	// Calculate the difference between encoded length and original code length
	diff, err := calculateEncodedDiff(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output the result
	fmt.Println("The difference between the total encoded characters and code characters is:", diff)
}
