package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type inputData struct {
	command     string
	category    string
	description string
}

func userInput(cat string, jd jsonData) inputData {
	in := inputData{}

	if isKommieCom(cat) {
		fmt.Printf("Cannot use '%s' as category\n", cat)
		return in
	}

	fmt.Print("Command           : ")
	com := readInput()
	if com == "" {
		fmt.Println("No command specified")
		return in
	}

	fmt.Print("Description (opt) : ")
	des := readInput()

	in.command = com
	in.category = cat
	in.description = des
	return in
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, " \n")
	return input
}

// Check if arg is a predefined command
func isKommieCom(arg string) bool {
	if arg == argAddCom || arg == argDelCom || arg == argModCom || arg == argDelCat || arg == argModCat {
		return true
	}
	return false
}
