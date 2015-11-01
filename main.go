package main

import (
	"fmt"
	"os"
)

const (
	argAddCom = "add"
	argDelCom = "del"
	argModCom = "mod"

	argDelCat = "delcat"
	argModCat = "modcat"
)

func main() {
	// Arguments
	args := os.Args[1:]
	// Argument count
	argc := len(args)

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
		cat := readInput()
		if cat == "" {
			fmt.Println("No category specified")
			break
		}

		argMatch(cat, args[0], jd)
	case 2:
		argMatch(args[0], args[1], jd)
	default:
		fmt.Println("Too many arguments")
	}
}

func argMatch(cat string, arg string, jd jsonData) {
	if arg == argAddCom {
		emptyInput := inputData{}
		in := userInput(cat, jd)
		if in != emptyInput {
			addCommand(in, &jd)
			writeJSONData(jd)
		}
		return
	}

	n := checkCategory(cat, jd)
	if n == -1 {
		fmt.Printf("No category named '%s'\n", cat)
		return
	}

	switch arg {
	case argDelCom:
		deleteCommand(n, &jd)
		writeJSONData(jd)
	case argModCom:
		modifyCommand(n, &jd)
		writeJSONData(jd)
	case argDelCat:
		deleteCategory(n, &jd)
		writeJSONData(jd)
	case argModCat:
		modifyCategory(n, &jd)
		writeJSONData(jd)
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
