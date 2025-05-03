package util

import (
	"compress/zlib"
	"fmt"
	"os"
)

func StoreObject(hash string, content string) error {
	objectPath := ".git/objects/" + hash[:2]
	err := os.MkdirAll(objectPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create the path for the object file: %w", err)
	}

	objectFilePath := objectPath + "/" + hash[2:40]
	file, err := os.Create(objectFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file the object file: %w", err)
	}

	zlibWriter := zlib.NewWriter(file)
	_, err = zlibWriter.Write([]byte(content))
	if err != nil {
		defer os.Remove(objectFilePath)
		return fmt.Errorf("failed to write into the object file: %w", err)
	}

	zlibWriter.Close()

	return nil
}
