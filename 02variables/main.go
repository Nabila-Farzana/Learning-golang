package main

import "fmt"

// public variable must contain first letter as capital
const LoginToken string = "dhvjdsj" //public

// jwtToken := 30000,  this does not work outside the method or like global variable

func main() {
	var username string
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//implicite type

	var website = "learningcodeline.in"
	fmt.Println(website)

	// no var style, can be changed into another type
	numberOfUser := 3000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
