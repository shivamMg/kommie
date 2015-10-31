package main

import (
	"bufio"
	"fmt"
	"os"
)

func modifyCommand(n int, jd *jsonData) {
	cats := &jd.Categories

	// Display all commands in nth category
	// Display serial-wise
	for i, com := range (*cats)[n].Commands {
		fmt.Printf("%d ", i+1)
		setColor()
		fmt.Println(com.Name)
		resetColor()
	}

	// Total commands in nth category
	l := len((*cats)[n].Commands)

	// Input serial number
	fmt.Print("Enter Serial Number: ")
	sno := 0
	_, err := fmt.Scanf("%d", &sno)
	// Check if sno is an integer and is in range (0, l]
	if err != nil || !(sno > 0 && sno <= l) {
		fmt.Println("Invalid Serial Number")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	com := readInput(reader)
	if com == "" {
		fmt.Println("No command specified")
		return
	}
	des := readInput(reader)

	(*cats)[n].Commands[sno-1].Name = com
	(*cats)[n].Commands[sno-1].Description = des
}

func modifyCategory(n int, jd *jsonData) {
	cats := &jd.Categories

	reader := bufio.NewReader(os.Stdin)
	cat := readInput(reader)
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
