package main

import (
	"fmt"
)

func AddCommand(input InputData, jsonData *JsonData) {
	categories := &jsonData.Categories
	for i, category := range *categories {
		// If category already exists
		if input.Category == category.Name {
			command := Command{
				Name:        input.Command,
				Description: input.Description,
			}
			(*categories)[i].Commands = append((*categories)[i].Commands, command)
			return
		}
	}
	// Category does not exist
	category := Category{
		Name: input.Category,
		Commands: []Command{
			{
				Name:        input.Command,
				Description: input.Description,
			},
		},
	}
	*categories = append(*categories, category)
}

func DeleteCommand(argCategory string, jsonData *JsonData) {
	categories := &jsonData.Categories

	n := -1
	for i, category := range *categories {
		if argCategory == category.Name {
			n = i
			break
		}
	}

	if n == -1 {
		fmt.Printf("No category named '%s'\n", argCategory)
		return
	} else {
		// Display all commands in nth category
		for i, command := range (*categories)[n].Commands {
			fmt.Printf("%d  %s\n", i+1, command.Name)
		}
	}

	fmt.Print("Enter Serial Number: ")
	serialno := 0
	_, err := fmt.Scanf("%d", &serialno)

	l := len((*categories)[n].Commands)
	// Check if serialno is an integer and is in range (0, l]
	// If yes, delete command at (serialno-1)
	if err == nil && serialno > 0 && serialno <= l {
		// Swap with last command
		(*categories)[n].Commands[serialno-1] = (*categories)[n].Commands[l-1]
		// Slice off last command
		(*categories)[n].Commands = (*categories)[n].Commands[:l-1]
	} else {
		fmt.Println("Invalid Serial Number")
	}
}
