package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")
	content := "This needs to go in a file - Nabila's habi jabi"

	file, err := os.Create("./myFile.txt") //Creation is a os operation
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close() // defer is recommended , because it needs to be closed after work done

	readFile("./myFile.txt")
}
func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // when we read file we read it not is string format , rather in byte format
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(dataByte)) // string() used to convert data byte in string

}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // panic will just shut down the program and show the error you are facing
	}
}
