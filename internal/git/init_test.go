package git

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, error := os.MkdirTemp("", "gogit-test")
	if error != nil {
		t.Fatalf("Failed to create temporary directory: %v", error)
	}
	defer os.RemoveAll(tempDir)

	// Save current directory
	currentDir, error := os.Getwd()
	if error != nil {
		t.Fatalf("Failed to get current directory: %v", error)
	}

	error = os.Chdir(tempDir)
	if error != nil {
		t.Fatalf("Failed to change to temporary directory: %v", error)
	}

	// Run the Init function
	Init()
}
