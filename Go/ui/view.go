package ui

import (
	"fmt"
)

func GetMenuOption() int {

	var userInput int

	fmt.Println("Main Menu")
	fmt.Println("1. List Records")
	fmt.Println("2. Create New Record")
	fmt.Println("3. View a Record")
	fmt.Println("4. Remove a Record")
	fmt.Println("5. Exit")

	fmt.Scanln(&userInput)

	return userInput
}

func GetIntInput(statement string) int {
	var userInput int
	fmt.Println(statement)

	fmt.Scanln(&userInput)
	return userInput
}
