package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //"01-02-2006 Monday" this time can be used to format. Though this seems crazy, But you have to remember this time string if you want to format this way.

	createDate := time.Date(2020, time.July, 21, 5, 30, 1, 21, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
