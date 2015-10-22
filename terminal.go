package main

import (
	"fmt"
	"os"
)

func CatchArgs() {
	args := os.Args[1:]
	argc := len(args)

	if argc == 0 {
		AddCommand()
	} else if argc == 1 {
		switch args[0] {
		case "show":
			jsonData := ReadJsonData()
			DisplayCategories(jsonData)
		case "delete":
			jsonData := ReadJsonData()
			DisplayCategories(jsonData)
			fmt.Println("Enter category name as second argument")
		}
	} else if argc == 2 {
		switch args[0] {
		case "show":
			DisplayCommands(args[1])
		case "delete":
			DeleteCommand(args[1])
		}
	}
}

func AddCommand() {
	jsonData := ReadJsonData()
	emptyUserInput := UserInput{}
	input := InputData(jsonData)
	if input != emptyUserInput {
		ModifyJsonData(input, &jsonData)
		WriteJsonData(jsonData)
	}
}

func DisplayCommands(argsCategory string) {
	jsonData := ReadJsonData()
	n := -1
	for i, category := range jsonData.Categories {
		if argsCategory == category.Name {
			n = i
			break
		}
	}
	if n == -1 {
		fmt.Printf("No category named '%s'\n", argsCategory)
	} else {
		for _, command := range jsonData.Categories[n].Commands {
			fmt.Printf("%s -- %s\n", command.Name, command.Description)
		}
	}
}

func DeleteCommand(argsCategory string) {
	jsonData := ReadJsonData()
	n := -1
	for i, category := range jsonData.Categories {
		if argsCategory == category.Name {
			n = i
			break
		}
	}
	if n == -1 {
		fmt.Printf("No category named '%s'\n", argsCategory)
	} else {
		for i, command := range jsonData.Categories[n].Commands {
			fmt.Printf("%d  %s\n", i+1, command.Name)
		}
	}
	l := len(jsonData.Categories[n].Commands)
	fmt.Print("Enter Serial Number: ")
	serialno := 0
	_, err := fmt.Scanf("%d", &serialno)
	if serialno > 0 && serialno <= l && err == nil {
		jsonData.Categories[n].Commands[serialno-1] = jsonData.Categories[n].Commands[l-1]
		jsonData.Categories[n].Commands = jsonData.Categories[n].Commands[:l-1]
		WriteJsonData(jsonData)
	} else {
		fmt.Println("Invalid Serial Number")
	}
}
