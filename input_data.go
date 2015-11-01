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

func readSerialNo(l int) int {
	fmt.Print("Enter Serial Number: ")
	sno := 0
	_, err := fmt.Scanf("%d", &sno)
	// if sno is not integer and is not in range (0, l]
	if err != nil || !(sno > 0 && sno <= l) {
		fmt.Println("Invalid Serial Number")
		return -1
	}
	return sno
}

// Check if arg is a predefined command
func isKommieCom(arg string) bool {
	if arg == argAddCom || arg == argDelCom || arg == argModCom || arg == argDelCat || arg == argModCat || arg == argExport {
		return true
	}
	return false
}
