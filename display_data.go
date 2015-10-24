package main

import (
	"fmt"
)

func displayCommands(argCategory string, jd jsonData) {
	// Search categories for argCategory
	// Get index
	n := -1
	for i, c := range jd.Categories {
		if argCategory == c.Name {
			n = i
			break
		}
	}

	if n == -1 {
		fmt.Printf("No category named '%s'\n", argCategory)
	} else {
		for _, com := range jd.Categories[n].Commands {
			fmt.Printf("%s -- %s\n", com.Name, com.Description)
		}
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
	i := 0
	for i < len(cats) {
		fmt.Print(cats[i])
		for k := 0; k <= maxCatLen-len(cats[i]); k++ {
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
