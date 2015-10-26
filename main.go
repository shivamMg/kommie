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
	const argDelete = "del"
	const argDeleteCategory = "delcat"

	switch argc {
	case 0:
		emptyInput := inputData{}
		in := userInput(jd)
		if in != emptyInput {
			addCommand(in, &jd)
			writeJSONData(jd)
		}
	case 1:
		displayCategories(jd)
		if args[0] == argShow {
			// Do nothing
		} else if args[0] == argDelete {
			fmt.Println("Enter category name as second argument")
		} else if args[0] == argDeleteCategory {
			fmt.Println("Enter category name as second argument")
		}
	case 2:
		if args[0] == argShow {
			displayCommands(args[1], jd)
		} else if args[0] == argDelete {
			deleteCommand(args[1], &jd)
			writeJSONData(jd)
		} else if args[0] == argDeleteCategory {
			deleteCategory(args[1], &jd)
			writeJSONData(jd)
		}
	}
}
