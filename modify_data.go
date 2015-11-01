package main

import (
	"fmt"
)

func modifyCommand(n int, jd *jsonData) {
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

	fmt.Print("New command name : ")
	com := readInput()
	if com == "" {
		fmt.Println("No command specified")
		return
	}
	fmt.Print("New description  : ")
	des := readInput()

	(*cats)[n].Commands[sno-1].Name = com
	(*cats)[n].Commands[sno-1].Description = des
}

func modifyCategory(n int, jd *jsonData) {
	cats := &jd.Categories

	fmt.Print("New category name : ")
	cat := readInput()
	if cat == "" {
		fmt.Println("No category specified")
		return
	}

	m := checkCategory(cat, *jd)
	if m != -1 {
		fmt.Printf("'%s' already exists\n", cat)
		return
	}

	(*cats)[n].Name = cat
}
