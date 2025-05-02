package git

import (
	"os"
	"testing"
)

func testSetup(t *testing.T) {
	// Mark the function as a helper
	t.Helper()

	// Create a temporary directory for testing
	testDir := t.TempDir()

	// Change to the test directory
	err := os.Chdir(testDir)
	if err != nil {
		t.Fatalf("Failed to change to test directory: %v", err)
	}
}

func TestInit(t *testing.T) {
	t.Run("succesful_init", func(t *testing.T) {
		// Setup
		testSetup(t)

		// Run
		isCreated := Init()

		// Verify
		if !isCreated {
			t.Fatalf("Failed to create .git directory")
		}

		_, err := os.Stat(".git")
		if os.IsNotExist(err) {
			t.Fatalf("Failed to create .git directory")
		}

		_, err = os.Stat(".git/objects")
		if os.IsNotExist(err) {
			t.Fatalf("Failed to create .git/objects directory")
		}

		_, err = os.Stat(".git/refs")
		if os.IsNotExist(err) {
			t.Fatalf("Failed to create .git/refs directory")
		}

		_, err = os.Stat(".git/HEAD")
		if os.IsNotExist(err) {
			t.Fatalf("Failed to create .git/HEAD file")
		}

		_, err = os.Stat(".git/config")
		if os.IsNotExist(err) {
			t.Fatalf("Failed to create .git/config file")
		}
	})

	t.Run("git_directory_already_exists", func(t *testing.T) {
		// Setup
		testSetup(t)

		// Run
		isCreated := Init()
		if !isCreated {
			t.Fatalf("Failed to create .git directory")
		}

		isCreatedTwice := Init()
		if isCreatedTwice {
			t.Fatalf(".git repository folder was initialized twice")
		}
	})
}
