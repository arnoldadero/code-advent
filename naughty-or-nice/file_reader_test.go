package main

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Create a temporary file with test data
	fileContent := "firstline\nsecondline\nthirdline\n"
	tmpFile, err := os.CreateTemp("", "test_input*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the file after test

	// Write content to the temp file
	if _, err := tmpFile.WriteString(fileContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Call ReadFile on the temp file
	lines, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Verify the results
	expected := []string{"firstline", "secondline", "thirdline"}
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, but got %d", len(expected), len(lines))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected line %d to be %q, but got %q", i, expected[i], line)
		}
	}
}
