package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file has a
first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and
print the first and last names found in each struct.

Submit your source code for the program, “read.go”.
*/

type Person struct {
	fname string
	lname string
}

func main() {

	var people []Person

	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Data Path? :")
	if scn.Scan() {
		path := scn.Text()
		f, err := os.Open(path)
		if err != nil {
			fmt.Println("Could not open file! Provided path:", path)
			return
		} else {
			defer f.Close()
			fileReader := bufio.NewScanner(f)
			for fileReader.Scan() {
				line := fileReader.Text()
				bits := strings.Split(line, " ")

				var fname string
				var lname string

				for idx, part := range bits {
					trunc := string([]rune(part)[:20])
					if idx == 0 {
						fname = trunc
					} else {
						lname = trunc
					}
				}

				people = append(people, Person{fname: fname, lname: lname})

			}

			for _, person := range people {
				fmt.Println("First name:", person.fname)
				fmt.Println("Last name:", person.lname, "\n---------------\n")
			}

		}
	}
}
