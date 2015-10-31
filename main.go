package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Arguments
	args := os.Args[1:]
	// Argument count
	argc := len(args)

	const (
		argAddCommand    = "add"
		argDeleteCommand = "del"
		argModifyCommand = "mod"

		argDeleteCategory = "delcat"
		argModifyCategory = "modcat"
	)

	jd := readJSONData()

	switch argc {
	case 0:
		displayCategories(jd)
	case 1:
		n := checkCategory(args[0], jd)
		// if arg[0] is a category
		if n != -1 {
			displayCommands(n, jd)
			break
		}

		displayCategories(jd)
		fmt.Print("Category          : ")
		reader := bufio.NewReader(os.Stdin)
		cat := readInput(reader)
		if cat == "" {
			fmt.Println("No category specified")
			break
		}

		if args[0] == argAddCommand {
			emptyInput := inputData{}
			in := userInput(cat, jd)
			if in != emptyInput {
				addCommand(in, &jd)
				writeJSONData(jd)
			}
			break
		}

		n = checkCategory(cat, jd)
		if n == -1 {
			fmt.Printf("No category named '%s'\n", cat)
			break
		}

		if args[0] == argDeleteCommand {
			deleteCommand(n, &jd)
			writeJSONData(jd)
		} else if args[0] == argDeleteCategory {
			deleteCategory(n, &jd)
			writeJSONData(jd)
		} else if args[0] == argModifyCommand {
			modifyCommand(n, &jd)
			writeJSONData(jd)
		} else if args[0] == argModifyCategory {
			modifyCategory(n, &jd)
			writeJSONData(jd)
		}
	case 2:
		n := checkCategory(args[0], jd)
		if n == -1 {
			fmt.Printf("No category named '%s'\n", args[0])
			break
		}

		if args[1] == argAddCommand {
			emptyInput := inputData{}
			in := userInput(args[0], jd)
			if in != emptyInput {
				addCommand(in, &jd)
				writeJSONData(jd)
			}
		} else if args[1] == argDeleteCommand {
			deleteCommand(n, &jd)
			writeJSONData(jd)
		} else if args[1] == argDeleteCategory {
			deleteCategory(n, &jd)
			writeJSONData(jd)
		} else if args[1] == argModifyCommand {
			modifyCommand(n, &jd)
			writeJSONData(jd)
		} else if args[1] == argModifyCategory {
			modifyCategory(n, &jd)
			writeJSONData(jd)
		}
	}
}

func checkCategory(category string, jd jsonData) int {
	n := -1
	for i, c := range jd.Categories {
		if category == c.Name {
			// Get index
			n = i
			break
		}
	}
	return n
}
