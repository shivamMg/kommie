package main

import (
	"os"
)

func CatchArgs() {
	args := os.Args[1:]
	argc := len(args)

	if argc == 0 {
		AddCommand()
	} /*else if argc == 1 {
		switch args[1] {
		case "show":
			DisplayCategories()
			DisplayCommands()
		case "delete":
			DisplayCategories()
			DeleteCommand()
		}
	} else if argc == 2 {
		switch args[1] {
		case "show":
			DisplayCommands()
		case "delete":
			DeleteCommand()
		}
	}*/
}

func AddCommand() {
	jsonData := ReadJsonData()
	emptyUserInput := UserInput{}
	input := InputData(*jsonData)
	if input != emptyUserInput {
		ModifyJsonData(input, jsonData)
		WriteJsonData(*jsonData)
		// DisplayData(jsonData)
	}
}
