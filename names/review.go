package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	// Name struct
	fname string
	lname string
}

// Check any method/function that returns an error (err)
// If true: apply panic()
// See: https://gobyexample.com/reading-files
func checkError(err error, message string) {
	defaultMessage := "Error encountered!\n"
	if message == "" {
		message = defaultMessage
	}
	if err != nil {
		fmt.Fprint(os.Stderr, message)
		// log.Fatal(err)
		panic(err)
	}
}

func readFile(filepath string) []string {
	var lines []string
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err) // log.Fatal ==> print error followed by panic
	}

	defer f.Close() // defer closing file until the end of function

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
		// fmt.Printf("\n%T\t", text)
		// fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err) // log.Fatal ==> print error followed by panic
	}
	//os.Exit(1) // temporarily placed to exit code
	return lines
}

func main() {
	/*
		Write a program which reads information from a file and represents
		it in a slice of structs. Assume that there is a text file which
		contains a series of names. Each line of the text file has a first
		name and a last name, in that order, separated by a single space
		on the line.

		Your program will define a name struct which has two fields, fname
		for the first name, and lname for the last name. Each field will be
		a string of size 20 (characters).

		Your program should prompt the user for the name of the text file.
		Your program will successively read each line of the text file and
		create a struct which contains the first and last names found in the
		file. Each struct created will be added to a slice, and after all
		lines have been read from the file, your program will have a slice
		containing one struct for each line in the file. After reading all
		lines from the file, your program should iterate through your slice
		of structs and print the first and last names found in each struct.

		Submit your source code for the program, “read.go”.
	*/

	var filepath string
	var names []Name

	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Enter filename: \t")
	// scanner.Scan()
	// filename = scanner.Text()
	filepath = getConsoleInput(scanner, "Enter filename: \t")
	fmt.Printf(">>> Filename: \t%s", filepath)

	// f, err := os.Open(filepath)
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, "Error opening file\n")
	// 	panic(err)
	// }
	// checkError(err, "Error opening file\n")

	// Read file as lines:
	//       a slice of strings, one string
	//       for each line.
	lines := readFile(filepath)

	// Split each line into a slice and
	//       extract the first two words.
	sep := " "
	for _, line := range lines {
		//fmt.Printf("%d. line: %s\n", idx, line)
		parts := strings.Split(line, sep)
		names = append(names, Name{fname: parts[0], lname: parts[1]})
	}

	// Iterate over the names and print
	//         fields: fname and lname.
	fmt.Println("")
	for idx, name := range names {
		fmt.Printf("%d. FirstName: %s, LastName: %s\n", idx, name.fname, name.lname)
	}

}

func getConsoleInput(scanner *bufio.Scanner, prompt string) string {
	var value string
	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(prompt)
	scanner.Scan()
	value = scanner.Text()
	return value
}

func printSlice(s []int) {
	fmt.Printf("Slice Info: \n\tlen=%d cap=%d slice=%v\n", len(s), cap(s), s)
}
