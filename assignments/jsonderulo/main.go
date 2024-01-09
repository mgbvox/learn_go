package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"unicode"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.

*/

func prompt(key string, scn *bufio.Scanner, hash *map[string]string) {
	if key == "" {
		return
	}
	keyRune := []rune(key)
	keyRune[0] = unicode.ToUpper(keyRune[0])
	prompt := string(keyRune) + ":"
	fmt.Println(prompt)
	if scn.Scan() {
		(*hash)[key] = scn.Text()
	}

}

func main() {
	dude := make(map[string]string)

	scn := bufio.NewScanner(os.Stdin)
	prompt("name", scn, &dude)
	prompt("address", scn, &dude)
	jdata, err := json.MarshalIndent(dude, "", "\t")
	if err != nil {
		fmt.Println("Error!")
	}
	fmt.Println(string(jdata))

}
