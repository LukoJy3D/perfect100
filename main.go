package main

import (
	"fmt"
	"os"
	utils "perfect100/.utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [<full game name>|all]")
		return
	}

	gameName := os.Args[1]

	utils.Achievements(gameName)
}
