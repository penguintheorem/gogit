package util

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func buildHeader(content []byte) string {
	byteSize := len(content)
	header := fmt.Sprintf("%s %d\x00", "blob", byteSize)
	return header
}

func GetObjectHash(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	header := buildHeader(content)
	combined := header + string(content)

	hash := sha1.New()
	hash.Write([]byte(combined))

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
