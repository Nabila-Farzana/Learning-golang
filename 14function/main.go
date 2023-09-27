package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()
	result := adder(3, 4)
	fmt.Println(result)

	proResult, message := proAdder(3, 5, 7, 9)
	fmt.Println(message, ":", proResult)
}

func adder(valOne int, valTwo int) int { // you have to mention the parameter types and the return type
	return valOne + valTwo
}
func greeter() {
	fmt.Println("Salam Golang!!")
}

func proAdder(values ...int) (int, string) { // values ...int means we will recieve many values of type intiger and the return value is an intiger and a string
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "The total amount is"

}
