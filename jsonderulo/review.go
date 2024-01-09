package main

import "encoding/json"
import "fmt"

func main() {

	fmt.Printf("Please enter your name and address.\n")

	dict := make(map[string]string)

	var input string
	fmt.Printf("Name: ")
	fmt.Scan(&input)
	dict["name"] = input

	fmt.Printf("Address: ")
	fmt.Scan(&input)
	dict["address"] = input
	b, err := json.Marshal(dict)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
