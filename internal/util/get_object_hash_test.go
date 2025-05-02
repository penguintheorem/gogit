package util

import (
	"os"
	"testing"
)

func testSetup(t *testing.T) string {
	// Mark the function as a helper
	t.Helper()

	// Create a temporary directory for testing
	testDir := t.TempDir()

	// Change to the test directory
	err := os.Chdir(testDir)
	if err != nil {
		t.Fatalf("Failed to change to test directory: %v", err)
	}

	return testDir
}

func createTestFile(t *testing.T, dir, content string) string {
	t.Helper()

	filePath := dir + "/test.txt"
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	return filePath
}

func TestInit(t *testing.T) {
	t.Run("input_file_path_does_not_exist", func(t *testing.T) {
		_, err := GetObjectHash("fake_path")
		if err == nil {
			t.Fatalf("The file path results to exist also if it is a fake path")
		}
	})
	t.Run("valid_input_file", func(t *testing.T) {
		testDir := testSetup(t)

		filePath := createTestFile(t, testDir, "hello world!")
		expectedHash := "bc7774a7b18deb1d7bd0212d34246a9b1260ae17"

		hash, err := GetObjectHash(filePath)
		if err != nil {
			t.Fatalf("GetObjectHash failed: %v", err)
		}

		if hash != expectedHash {
			t.Errorf("GetObjectHash = %v, want %v", hash, expectedHash)
		}
	})

	t.Run("input_file_is_empty", func(t *testing.T) {
		testDir := testSetup(t)

		filePath := createTestFile(t, testDir, "")
		expectedHash := "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"

		hash, err := GetObjectHash(filePath)
		if err != nil {
			t.Fatalf("GetObjectHash failed: %v", err)
		}

		if hash != expectedHash {
			t.Errorf("GetObjectHash = %v, want %v", hash, expectedHash)
		}
	})

}
