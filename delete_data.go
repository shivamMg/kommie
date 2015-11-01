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

	displayCommands(n, true, *jd)
	sno := readSerialNo(l)
	if sno == -1 {
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

	fmt.Printf("'%s' contains %d command(s)\n", (*cats)[n].Name, l)

	ch := "N"
	fmt.Printf("You sure you want to delete '%s'? (y/N): ", (*cats)[n].Name)
	_, err := fmt.Scanf("%s", &ch)
	if err != nil || (ch != "y" && ch != "Y") {
		return
	}

	// Swap with last category
	(*cats)[n] = (*cats)[l-1]
	// Slice off last category
	*cats = (*cats)[:l-1]
}
