package main

import "fmt"

func main() {
	fmt.Println("Method in golang")

	nabila := User{"Nabila", "nabila@go.dev", true, 21}
	fmt.Println(nabila)
	fmt.Printf("Nabila's details are: %+v\n", nabila)

	fmt.Printf("Name is %v and email is %v.\n", nabila.Name, nabila.Email)
	nabila.GetStatus()
	nabila.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int // this property is not exportable as it has small first letter
}

//Method
func (u User) GetStatus() { // GetStatus must have capital first letter if I want it to export
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() { // This User is actually a copy of the User structure, not the main structure refference. So Main User won't be changed
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
