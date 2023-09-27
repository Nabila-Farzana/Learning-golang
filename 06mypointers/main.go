package main

import "fmt"

// why pointers? - Sometimes we pass a variable into various functionalities and thus a copy of this variable can be cretaed and a irregularity can be caused. To avoid this irregularity pointer concept came.Pointer carries the memry address of the variable, So whenever you pass a pointer you actually pass the memory location of the variable. So no chance for duplication.
func main() {
	fmt.Println("Welcome to a class on pointer")
	// var ptr *int // a pointer storing intiger type value
	// fmt.Println("Value of pointer is ", ptr) // default value is nil

	myNumber := 23
	var ptr = &myNumber                               //Reference means &
	fmt.Println("Value of actual pointer is  ", ptr)  // output the memory address
	fmt.Println("Value of actual pointer is  ", *ptr) // output 23
	// Pointer make sure that if we do any operation on the value , it actually performs the opration on that value , not on the copy of that value
	*ptr = *ptr * 2
	fmt.Println("New value is ", myNumber) // output is 46
}
