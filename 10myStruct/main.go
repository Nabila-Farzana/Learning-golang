package main

import "fmt"

func main() {
	// no inheritance in golang; No super or parent
	fmt.Println("Structs in golang")

	nabila := User{"Nabila", "nabila@go.dev", true, 21}
	fmt.Println(nabila)                               // output: {Nabila nabila@go.dev true 21}
	fmt.Printf("Nabila's details are: %+v\n", nabila) // output : {Name:Nabila Email:nabila@go.dev Status:true Age:21}. here %+v is for more details

	fmt.Printf("Name is %v and email is %v.\n", nabila.Name, nabila.Email)
}

// defining a struct
type User struct { // This is mostly lika a template or class , So struct name must start with capital letter
	Name   string // All the field should have capital first letter
	Email  string
	Status bool
	Age    int
}
