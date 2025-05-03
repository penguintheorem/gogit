package git

import (
	"fmt"
	"gogit/internal/util"
	"os"
)

func Add(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	header := util.BuildHeader(content)
	combined := header + string(content)
	hash := util.GetObjectHash(combined)

	err = util.StoreObject(hash, combined)
	if err != nil {
		return "", fmt.Errorf("failed to store the file: %w", err)
	}

	return "", nil
}
