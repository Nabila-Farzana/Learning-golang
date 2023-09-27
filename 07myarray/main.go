package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitlist [4]string // array capacity is required

	fruitlist[0] = "apple"
	fruitlist[1] = "Tomato"
	fruitlist[3] = "peach"

	fmt.Println("Fruit list is: ", fruitlist) // As third element is missing, it will keep a small space in output for the indication of the missing element

	fmt.Println("Fruit list length is: ", len(fruitlist)) // output will be 4 , no matter what how many elements are in the list, it will always show 4 as I declared it at the time of initializing

	var veglist = [3]string{"potato", "beans", "mushroom"} // decalre and initialize
	fmt.Println("Vegy list is: ", veglist)
}
