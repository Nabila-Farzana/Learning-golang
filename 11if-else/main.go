package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 10
	var result string
	if loginCount < 10 { //  There is no option to put this curly brace to the next line in golang
		result = "Regulate user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 loginCount"
	}
	fmt.Println(result)

	if 9%2 == 0 { // we can write logic this way
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Web req handling
	if num := 3; num < 10 { // while working with web req and check the value on the go
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil { for checking error

	// }
}
