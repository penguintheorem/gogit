package util

import (
	"testing"
)

func TestBuildHeader(t *testing.T) {
	t.Run("encode_non_empty_content_into_header", func(t *testing.T) {
		header := BuildHeader([]byte("hello"))
		expectedHeader := "blob 5\x00"
		if header != expectedHeader {
			t.Errorf("BuildHeader = %v, want %v", header, expectedHeader)
		}
	})

	t.Run("encode_empty_content_into_header", func(t *testing.T) {
		header := BuildHeader([]byte{})
		expectedHeader := "blob 0\x00"
		if header != expectedHeader {
			t.Errorf("BuildHeader = %v, want %v", header, expectedHeader)
		}
	})
}
