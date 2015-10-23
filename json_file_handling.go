package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var JsonFileName = "./data.json"

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

func ReadJsonData() JsonData {
	jsonData := JsonData{}
	content, err := ioutil.ReadFile(JsonFileName)
	// If file already exists
	if err == nil {
		err = json.Unmarshal(content, &jsonData)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	return jsonData
}

func WriteJsonData(jsonData JsonData) {
	content, err := json.Marshal(jsonData)
	err = ioutil.WriteFile(JsonFileName, content, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
