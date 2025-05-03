package cli

import (
	"fmt"
	"gogit/internal/git"
	"os"
)

func Add(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("Error: File does not exist")
		return
	}

	_, err = git.Add(filePath)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
