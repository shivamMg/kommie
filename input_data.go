package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputData struct {
	Command     string
	Category    string
	Description string
}

func UserInput(jsonData JsonData) InputData {
	input := InputData{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Command           : ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimRight(command, " \n")
	if command == "" {
		fmt.Println("No command specified")
		return input
	}

	DisplayCategories(jsonData)
	fmt.Print("Category          : ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimRight(category, " \n")
	if category == "" {
		fmt.Println("No category specified")
		return input
	}

	fmt.Print("Description (opt) : ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimRight(description, " \n")

	input.Command = command
	input.Category = category
	input.Description = description
	return input
}
