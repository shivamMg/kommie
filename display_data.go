package main

import (
	"fmt"
)

func DisplayCommands(argCategory string, jsonData JsonData) {
	// Get category index for argCategory
	n := -1
	for i, category := range jsonData.Categories {
		if argCategory == category.Name {
			n = i
			break
		}
	}

	if n == -1 {
		fmt.Printf("No category named '%s'\n", argCategory)
	} else {
		for _, command := range jsonData.Categories[n].Commands {
			fmt.Printf("%s -- %s\n", command.Name, command.Description)
		}
	}
}

func DisplayCategories(jsonData JsonData) {
	// No categories defined
	if jsonData.Categories == nil {
		return
	}

	var categories []string
	var maxCategoryLen int = 0

	// Extract all category names
	for _, category := range jsonData.Categories {
		categories = append(categories, category.Name)
		// Length of Category with Max Length
		if maxCategoryLen < len(category.Name) {
			maxCategoryLen = len(category.Name)
		}
	}

	// Display categories in square matrix
	// Calculate required order for square matrix
	n := 1
	for len(categories) > n*n {
		n++
	}
	// Display categories
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
	// If number of categories less than size of matrix
	if i < n*n {
		fmt.Println()
	}
}
