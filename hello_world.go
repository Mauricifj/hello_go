package main

import (
	"fmt"
	"os"
)

func main() {
	welcome()

	showMenu()

	option := readOption()

	switch option {
	case 1:
		fmt.Println("Hello, World")
	case 2:
		showCustomHelloWorld()
	case 0:
		fmt.Println("Process finishing...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
		os.Exit(-1)
	}
}

func showCustomHelloWorld() {
	fmt.Print("Type your name, please: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello,", name)
}

func showMenu() {
	fmt.Println("--- MENU ---")
	fmt.Println("1 - Show Hello World")
	fmt.Println("2 - Show custom Hello World")
	fmt.Println("0 - Quit")
}

func welcome() {
	version := "1.0"
	fmt.Println("*** Program HelloWorld - Version", version, "***")
}

func readOption() int {
	fmt.Print("Choose one option, please: ")
	var option int
	fmt.Scan(&option)
	return option
}
