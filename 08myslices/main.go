package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	//Slices are actually under the hood arrays and these arrays when have some abstruction layers and some features they are called slices

	var fruitList = []string{"Apple", "Tomato", "Peach"}         // decalration and initialization
	fmt.Printf("The type of  this fruitList is %T\n", fruitList) //Checking the type. The type is []string

	fruitList = append(fruitList, "Mango", "Banana") // add elements
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // To slice up the Slice
	fmt.Println(fruitList)            // outputs without the first element as the first element is sliced out

	fruitList = append(fruitList[1:3]) // outputs [Tomato, Peach]. It won't count the 4th element as slice is inclusive

	highScores := make([]int, 4) // here []int = what type of data I want to have, 4 = what size of data I want to have

	highScores[0] = 234
	highScores[1] = 285
	highScores[2] = 336
	highScores[3] = 297
	// highScores[5] = 934, this will occur an error , because I fixed the size before

	//    But if i use append it will allow to add elements. Because it will reallocate the memory inspite of using the default size.
	highScores = append(highScores, 234, 678, 346, 335475)
	fmt.Println(highScores)

	sort.Ints(highScores) // will sor the slice into increasing order
	fmt.Println(highScores)
	sort.IntsAreSorted(highScores) // will return a boolean value depends on the slice is sorted or not

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2                                       // I want to remove the second index
	courses = append(courses[:index], courses[index+1:]...) // courses[:index] means the value from 0 index to value before index
	//courses[index+1:] means from value after the index to the last value

	fmt.Println(courses)
}
