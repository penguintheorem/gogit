package main

import (
	"fmt"
	"os"

	"gogit/internal/cli"
)

func main() {
	// os.Args[0] is the executable path
	// os.Args[1] is the real first argument
	if len(os.Args) < 2 {
		fmt.Println(cli.PrintHelp())
		return 
	}

	switch os.Args[1] {
		case "help", "--help":
			fmt.Println(cli.PrintHelp())
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			fmt.Println(cli.PrintHelp())
	}
}