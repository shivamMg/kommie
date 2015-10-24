package main

import (
	"fmt"
	"os"
)

func main() {
	// Arguments
	args := os.Args[1:]
	// Argument count
	argc := len(args)

	jd := readJSONData()

	const argShow = "show"
	const argDelete = "delete"

	switch argc {
	case 0:
		emptyInput := inputData{}
		in := userInput(jd)
		if in != emptyInput {
			addCommand(in, &jd)
			writeJSONData(jd)
		}
	case 1:
		if args[0] == argShow {
			displayCategories(jd)
		} else if args[0] == argDelete {
			displayCategories(jd)
			fmt.Println("Enter category name as second argument")
		}
	case 2:
		if args[0] == argShow {
			displayCommands(args[1], jd)
		} else if args[0] == argDelete {
			deleteCommand(args[1], &jd)
			writeJSONData(jd)
		}
	}
}
