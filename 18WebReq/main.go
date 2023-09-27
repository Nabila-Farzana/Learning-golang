package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println(" Hello to web request")
	response, err := http.Get(url)

	if err != nil {
		panic(err) // panic will just shut down the program and show the error you are facing
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close this connection. Otherwise it will be open always

	dataBytes, err := ioutil.ReadAll(response.Body) //ioutil is used for most of the read operation

	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}
