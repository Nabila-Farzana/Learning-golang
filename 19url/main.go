package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dgjksdgj"

func main() {
	fmt.Println("Hello to URL")
	fmt.Println(myurl)

	// parsing the url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()                                 // Directly access the query parameters as key value pair
	fmt.Printf("The type of query params are: %T\n", qparams) // output : url.Values (Basically a key value pair)

	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	partsOfUrl := &url.URL{ // It has to be used as refrence not as a copy
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=nabila",
	}

	anotherURL := partsOfUrl.String() // we can use String() like this as well
	fmt.Println(anotherURL)
}
