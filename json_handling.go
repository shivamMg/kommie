package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Command struct {
	Name        string
	Description string
}

type Category struct {
	Name     string
	Commands []Command
}

type JsonData struct {
	Categories []Category
}

func ReadJsonData() *JsonData {
	data := JsonData{}
	content, err := ioutil.ReadFile(JsonFileName)
	if err == nil {
		// If file exists
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	return &data
}

func ModifyJsonData(input UserInput, data *JsonData) {
	for i, category := range data.Categories {
		if input.Category == category.Name {
			// Category already exists
			command := Command{
				Name:        input.Command,
				Description: input.Description,
			}
			data.Categories[i].Commands = append(data.Categories[i].Commands, command)
			fmt.Println(data)
			return
		}
	}
	// Category does not exist
	category := Category{
		Name: input.Category,
		Commands: []Command{
			{
				Name:        input.Command,
				Description: input.Description,
			},
		},
	}
	data.Categories = append(data.Categories, category)
}

func WriteJsonData(data JsonData) {
	content, err := json.Marshal(data)
	err = ioutil.WriteFile(JsonFileName, content, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
