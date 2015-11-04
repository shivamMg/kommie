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
	fmt.Printf("New category '%s' created\n", in.category)
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
