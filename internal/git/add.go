package git

import (
	"fmt"
	"gogit/internal/util"
	"log"
)

func Add(filePath string) {
	hash, err := util.GetObjectHash(filePath)
	if err != nil {
		log.Printf("Failed to get object hash: %v", err)
		return
	}

	fmt.Printf("Here is your hash: %s", hash)
}
