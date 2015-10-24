package main

import (
	"fmt"
)

func addCommand(in inputData, jd *jsonData) {
	cats := &jd.Categories

	for i, c := range *cats {
		// If category already exists
		if in.category == c.Name {
			command := command{
				Name:        in.command,
				Description: in.description,
			}
			(*cats)[i].Commands = append((*cats)[i].Commands, command)
			return
		}
	}

	// Category does not exist
	category := category{
		Name: in.category,
		Commands: []command{
			{
				Name:        in.command,
				Description: in.description,
			},
		},
	}
	*cats = append(*cats, category)
}

func deleteCommand(argCategory string, jd *jsonData) {
	cats := &jd.Categories

	// Check if argCategory exists
	// Get index
	n := -1
	for i, c := range *cats {
		if argCategory == c.Name {
			n = i
			break
		}
	}

	if n == -1 {
		fmt.Printf("No category named '%s'\n", argCategory)
		return
	}

	// Display all commands in nth category
	// Display serial-wise
	for i, com := range (*cats)[n].Commands {
		fmt.Printf("%d  %s\n", i+1, com.Name)
	}
	// Input serial number
	fmt.Print("Enter Serial Number: ")
	sno := 0
	_, err := fmt.Scanf("%d", &sno)

	// Number of commands in nth category
	l := len((*cats)[n].Commands)

	// Check if sno is an integer and is in range (0, l]
	// If yes, delete command at (sno-1)
	if err == nil && sno > 0 && sno <= l {
		// Swap with last command
		(*cats)[n].Commands[sno-1] = (*cats)[n].Commands[l-1]
		// Slice off last command
		(*cats)[n].Commands = (*cats)[n].Commands[:l-1]
	} else {
		fmt.Println("Invalid Serial Number")
	}
}
