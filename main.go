package main

func main() {
	jsonFile := "./data.json"
	jsonData := ReadJsonData(jsonFile)
	emptyUserInput := UserInput{}
	kommie := InputData()
	if kommie != emptyUserInput {
		ModifyData(kommie, jsonData)
		WriteJsonData(jsonFile, *jsonData)
		// DisplayData(jsonData)
	}
}
