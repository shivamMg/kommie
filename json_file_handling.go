package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var jsonFileName = "/.data.json"
var jsonFilePath = "/home/" + os.Getenv("USER") + jsonFileName

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

	content, err := ioutil.ReadFile(jsonFilePath)
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

	err = ioutil.WriteFile(jsonFilePath, content, 0666)
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
