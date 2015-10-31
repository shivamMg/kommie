package main

import (
	"fmt"
)

func deleteCommand(n int, jd *jsonData) {
	cats := &jd.Categories

	// Total commands inside category
	l := len((*cats)[n].Commands)
	if l == 0 {
		fmt.Printf("No commands in '%s'\n", (*cats)[n].Name)
		return
	}

	// Display all commands in nth category
	// Display serial-wise
	for i, com := range (*cats)[n].Commands {
		fmt.Printf("%d ", i+1)
		setColor()
		fmt.Println(com.Name)
		resetColor()
	}

	// Input serial number
	fmt.Print("Enter Serial Number: ")
	sno := 0
	_, err := fmt.Scanf("%d", &sno)
	// Check if sno is an integer and is in range (0, l]
	if err != nil || !(sno > 0 && sno <= l) {
		fmt.Println("Invalid Serial Number")
		return
	}

	// Swap with last command
	(*cats)[n].Commands[sno-1] = (*cats)[n].Commands[l-1]
	// Slice off last command
	(*cats)[n].Commands = (*cats)[n].Commands[:l-1]
}

func deleteCategory(n int, jd *jsonData) {
	cats := &jd.Categories

	// Total commands inside category
	l := len((*cats)[n].Commands)

	fmt.Printf("'%s' contains %d commands\n", (*cats)[n].Name, l)

	ch := "N"
	fmt.Printf("Are you sure, you want to delete '%s'? (y/N): ", (*cats)[n].Name)
	_, err := fmt.Scanf("%s", &ch)
	if err != nil || (ch != "y" && ch != "Y") {
		return
	}

	// Total categories
	l = len(*cats)
	if l == 0 {
		return
	}
	// Swap with last category
	(*cats)[n] = (*cats)[l-1]
	// Slice off last category
	*cats = (*cats)[:l-1]
}
