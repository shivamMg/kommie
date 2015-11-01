package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var jsonFileName = "./data.json"

type command struct {
	Name        string `json:"name"`
	Description string `json:"use"`
}

type category struct {
	Name     string    `json:"name"`
	Commands []command `json:"commands"`
}

type jsonData struct {
	Categories []category `json:"category"`
}

func readJSONData() jsonData {
	jd := jsonData{}

	content, err := ioutil.ReadFile(jsonFileName)
	// If file already exists
	if err == nil {
		err = json.Unmarshal(content, &jd)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	return jd
}

func writeJSONData(jd jsonData) {
	content, err := json.Marshal(jd)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ioutil.WriteFile(jsonFileName, content, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func exportJSONData(jd jsonData) {
	content, err := json.MarshalIndent(jd, "", "\t")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(content))
}
