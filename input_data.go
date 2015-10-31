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
	reader := bufio.NewReader(os.Stdin)

	if cat == "add" || cat == "del" || cat == "mod" || cat == "delcat" || cat == "modcat" {
		fmt.Printf("Cannot use '%s' as category\n", cat)
		return in
	}

	fmt.Print("Command           : ")
	com := readInput(reader)
	if com == "" {
		fmt.Println("No command specified")
		return in
	}

	fmt.Print("Description (opt) : ")
	des := readInput(reader)

	in.command = com
	in.category = cat
	in.description = des
	return in
}

func readInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	input = strings.TrimRight(input, " \n")
	return input
}
