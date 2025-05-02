package cli

import (
	"fmt"
	"os"
)

func Add(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("Error: File does not exist")
		return
	}

}
