package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	argc := len(args)
	jsonData := ReadJsonData()

	// Argument count
	switch argc {
	case 0:
		emptyInput := InputData{}
		input := UserInput(jsonData)
		if input != emptyInput {
			AddCommand(input, &jsonData)
			WriteJsonData(jsonData)
		}
	case 1:
		if args[0] == "show" {
			DisplayCategories(jsonData)
		} else if args[0] == "delete" {
			DisplayCategories(jsonData)
			fmt.Println("Enter category name as second argument")
		}
	case 2:
		if args[0] == "show" {
			DisplayCommands(args[1], jsonData)
		} else if args[0] == "delete" {
			DeleteCommand(args[1], &jsonData)
			WriteJsonData(jsonData)
		}
	}
}
