package main

import (
	"bufio"
	"fmt"
	"os"
)

// comma ok syntax
func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza:")

	// comma ok syntax or err ok syntax
	input, _ := reader.ReadString('\n') // If I need to catch error , I can write err instead of _
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("The type of  this rating is %T\n", input)
}
