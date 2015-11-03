package main

import (
	"fmt"
)

func setBoldColor() {
	// Yellow-Bold text
	fmt.Print("\x1b[33;1m")
}

func setColor() {
	// Yellow text
	fmt.Print("\x1b[33m")
}

func resetColor() {
	// Default text
	fmt.Print("\x1b[0m")
}

func displayCommands(n int, serialize bool, jd jsonData) {
	// List of all commands
	coms := jd.Categories[n].Commands
	// Length of command with max length
	var maxComLen int

	// Calculate maxComLen
	for _, c := range coms {
		if maxComLen < len(c.Name) {
			maxComLen = len(c.Name)
		}
	}

	// Display commands
	for i, c := range coms {
		if serialize {
			fmt.Printf("%d   ", i+1)
		}
		setColor()
		fmt.Print(c.Name)
		resetColor()
		for j := 0; j < maxComLen-len(c.Name); j++ {
			fmt.Print(" ")
		}
		fmt.Printf("   %s\n", c.Description)
	}
}

func displayCategories(jd jsonData) {
	// If no categories
	if jd.Categories == nil {
		return
	}

	// List of all categories
	var cats []string
	// Length of category with max length
	var maxCatLen int

	// Extract all category names
	for _, c := range jd.Categories {
		cats = append(cats, c.Name)
		if maxCatLen < len(c.Name) {
			maxCatLen = len(c.Name)
		}
	}

	// Display categories in square matrix
	// Calculate required order for square matrix
	n := 1
	for len(cats) > n*n {
		n++
	}

	// Display categories
	setBoldColor()
	for i, c := range cats {
		fmt.Print(c)
		for j := 0; j <= maxCatLen-len(c); j++ {
			fmt.Print(" ")
		}
		if (i+1)%n == 0 {
			fmt.Println()
		}
	}
	resetColor()

	if len(cats)%n != 0 {
		fmt.Println()
	}
}

func displayHelp() {
	h := `A valid command is of the form:

	kommie <category> <operation>

<operation> can be one of the following:

	add     Add command to <category>
	del     Delete command from <category>
	mod     Modify command in <category>
	delcat  Delete <category>
	modcat  Modify <category>

If <operation> is not specified, commands inside <category> are displayed.
If <category> is not specified, all categories are displayed and user is asked for input. The corresponding <operation> is then performed.
If both are not specified, all categories are displayed.

Besides these, there are two more commands:

	kommie export

to export stored json data.

	kommie help

to display help.`
	fmt.Println(h)
}
