package cli

import (
	"fmt"
	"gogit/internal/git"
	"os"
)

func Init() {
	isCreated := git.Init()
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: Failed to get current directory")
		return
	}

	if isCreated {
		fmt.Printf("Initialized empty Git repository in %s/.git/\n", currentDirectory)
	}
}
