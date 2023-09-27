package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 //range  is  always non inclusive. So rand.Intn(6) will generate random number between 1 to 5. But we want a random number between 1 to 6, hence +1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")

	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // automatically run the next next switch case as well
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 2 spot and roll dice again")
	default:
		fmt.Println("Waht was that!")
	}
}
