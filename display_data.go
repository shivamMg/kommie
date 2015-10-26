package main

import (
	"fmt"
)

func setColor() {
	// Yellow text color
	fmt.Print("\x1b[33;1m")
}

func resetColor() {
	// Default text color
	fmt.Print("\x1b[0m")
}

func displayCommands(argCategory string, jd jsonData) {
	// Search categories for argCategory
	n := -1
	for i, c := range jd.Categories {
		if argCategory == c.Name {
			// Get index
			n = i
			break
		}
	}

	if n == -1 {
		fmt.Printf("No category named '%s'\n", argCategory)
		return
	}

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
	for _, c := range coms {
		setColor()
		fmt.Print(c.Name)
		resetColor()
		for i := 0; i <= maxComLen-len(c.Name); i++ {
			fmt.Print(" ")
		}
		fmt.Println(c.Description)
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
	setColor()
	for i, c := range cats {
		fmt.Print(c)
		for j := 0; j <= maxCatLen-len(c); j++ {
			fmt.Print(" ")
		}
		if i == len(cats)-1 {
			break
		}
		if (i+1)%n == 0 {
			fmt.Println()
		}
	}
	resetColor()

	// If number of categories less than size of matrix
	if len(cats) < n*n {
		fmt.Println()
	}
}
