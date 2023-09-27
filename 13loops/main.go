package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")

	days := []string{"Sunday", "TuesDay", "Wednesday", "ThurseDay", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// 	for i := range days { // range will automatically loop through the whole array
	// 		fmt.Println(days[i])
	// 	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
	rougueValue := 1

	for rougueValue < 10 { // Like while loop

		if rougueValue == 2 {
			goto sayHi
		}
		// if rougueValue == 5 { // for breaking the loop
		// 	break
		// }
		if rougueValue == 5 { // for continue  the loop
			rougueValue++
			continue
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++

	}

	// goto statement
sayHi:
	println("Hello, come to the goto statement!!")

}
