/*
Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string:
	* starts with the character ‘i’,
	* ends with the character ‘n’,
	* contains the character ‘a’.

The program should print “Not Found!” otherwise.

The program should not be case-sensitive, so it does not matter if the characters are upper-case
or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func HasIan(str string) bool {
	chars := []rune(strings.ToLower(str))
	valid := false
	for idx, c := range chars {
		switch {
		case idx == 0 && c != 'i':
			// start is not 'i'
			return false
		case idx == len(chars)-1 && c != 'n':
			// end is not 'n'
			return false
		case idx != 0 && idx != len(chars)-1:
			// somewhere in the middle...
			if c == 'a' {
				// ...a char was 'a'
				valid = true
			}
		}
	}

	return valid

}

//
//func main() {
//	var inp string
//	fmt.Println("Input a string: ")
//	_, err := fmt.Scan(&inp)
//	if err != nil {
//		fmt.Println("Bad input!")
//	}
//	if HasIan(inp) {
//		fmt.Println("Found!")
//	} else {
//		fmt.Println("Not Found!")
//	}
//}

func review() {
	fmt.Println("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Convert the string to lower case for case-insensitive comparison
		lowerInput := strings.ToLower(input)

		if strings.HasPrefix(lowerInput, "i") &&
			strings.HasSuffix(lowerInput, "n") &&
			strings.Contains(lowerInput, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}

func main() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}
