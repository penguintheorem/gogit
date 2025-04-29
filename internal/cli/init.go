package cli

import (
	"fmt"
	"gogit/internal/git"
	"os"
)

func Init() {
	git.Init()

	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: Failed to get current directory")
		return
	}
	fmt.Printf("Initialized empty Git repository in %s/.git/\n", currentDirectory)
}
