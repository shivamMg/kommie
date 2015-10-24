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

func userInput(jd jsonData) inputData {
	in := inputData{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Command           : ")
	com, _ := reader.ReadString('\n')
	com = strings.TrimRight(com, " \n")
	if com == "" {
		fmt.Println("No command specified")
		return in
	}

	displayCategories(jd)
	fmt.Print("Category          : ")
	cat, _ := reader.ReadString('\n')
	cat = strings.TrimRight(cat, " \n")
	if cat == "" {
		fmt.Println("No category specified")
		return in
	}

	fmt.Print("Description (opt) : ")
	des, _ := reader.ReadString('\n')
	des = strings.TrimRight(des, " \n")

	in.command = com
	in.category = cat
	in.description = des
	return in
}
