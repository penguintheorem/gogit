package util

import (
	"fmt"
)

func BuildHeader(content []byte) string {
	byteSize := len(content)
	header := fmt.Sprintf("%s %d\x00", "blob", byteSize)

	return header
}
