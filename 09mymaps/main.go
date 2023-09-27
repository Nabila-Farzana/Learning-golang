package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js shorts for: ", languages["js"])

	// Deleting item from map
	delete(languages, "rb")
	fmt.Println(languages)

	//Loop through the map
	for key, value := range languages { //for _, value := range languages -> we can write this line this way, if we want to skip the key
		fmt.Printf("For key %v, Value is %v\n", key, value) // %v is for value
	}
}
