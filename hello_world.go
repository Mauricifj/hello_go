package main
import "fmt"

func main() {
	version := "1.0"
	fmt.Println("*** Program HelloWorld - Version", version, "***")
	fmt.Println("--- MENU ---")
	fmt.Println("1 - Show Hello World")
	fmt.Println("2 - Show custom Hello World")
	fmt.Println("0 - Quit")

	fmt.Print("Choose one option, please: ")
	option := -1
	fmt.Scan(&option)
	if option == 1 {
		fmt.Println("Hello, World")
	} else if option == 2 {
		fmt.Print("Type your name, please: ")
		var name string
		fmt.Scan(&name)
		fmt.Println("Hello,", name)
	} else if option == 0 {
		fmt.Println("Process finishing...")
	} else {
		fmt.Println("Invalid option")
	}
	fmt.Println("Process ended successfully")
}
