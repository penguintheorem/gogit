package util

import (
	"crypto/sha1"
	"fmt"
)

func GetObjectHash(content string) string {
	hash := sha1.New()
	hash.Write([]byte(content))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
