package main

import (
	"fmt"
	"os"
	utils "perfect100/.utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [achievements|leaderboards]")
		return
	}

	command := os.Args[1]

	switch command {
	case "achievements":
		utils.Achievements()
	//Implement later
	//case "leaderboards":
	//	utils.Leaderboards()
	default:
		fmt.Println("Invalid command. Use 'achievements' or 'leaderboards'.")
	}
}
