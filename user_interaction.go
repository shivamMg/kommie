package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	Command     string
	Category    string
	Description string
}

func InputData(jsonData JsonData) UserInput {
	userInput := UserInput{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Command           : ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, " \n")
	if command == "" {
		fmt.Println("No command specified")
		return userInput
	}

	DisplayCategories(jsonData)
	fmt.Print("Category          : ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimRight(category, " \n")
	if category == "" {
		fmt.Println("No category specified")
		return userInput
	}

	fmt.Print("Description (opt) : ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimRight(description, " \n")

	userInput.Command = command
	userInput.Category = category
	userInput.Description = description
	return userInput
}

func DisplayData(data *JsonData) {
	for _, category := range data.Categories {
		fmt.Println(category.Name)
		for _, command := range category.Commands {
			fmt.Print(" " + command.Name)
			fmt.Print(" -- " + command.Description)
			fmt.Println()
		}
	}
}

func DisplayCategories(jsonData JsonData) {
	var categories []string
	var maxCategoryLen int = 0
	for _, category := range jsonData.Categories {
		categories = append(categories, category.Name)
		// Length of Category with Max Length
		if len(category.Name) > maxCategoryLen {
			maxCategoryLen = len(category.Name)
		}
	}
	n := 1
	for len(categories) > n*n {
		n++
	}
	i := 0
	for i < len(categories) {
		fmt.Print(categories[i])
		for k := 0; k <= maxCategoryLen-len(categories[i]); k++ {
			fmt.Print(" ")
		}
		if (i+1)%n == 0 {
			fmt.Println()
		}
		i++
	}
	if i < n*n {
		fmt.Println()
	}
}
