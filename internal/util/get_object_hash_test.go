package util

import (
	"testing"
)

func TestGetObjectHash(t *testing.T) {
	t.Run("valid_text_content", func(t *testing.T) {
		content := "blob 12\x00" + "hello world!"
		expectedHash := "bc7774a7b18deb1d7bd0212d34246a9b1260ae17"

		hash := GetObjectHash(content)
		if hash != expectedHash {
			t.Errorf("GetObjectHash = %v, want %v", hash, expectedHash)
		}
	})

	t.Run("empty_content", func(t *testing.T) {
		content := "blob 0\x00" + string("")
		expectedHash := "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391"

		hash := GetObjectHash(content)
		if hash != expectedHash {
			t.Errorf("GetObjectHash = %v, want %v", hash, expectedHash)
		}
	})
}
