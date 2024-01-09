package main

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.

[*] The program should be written as a loop.

[*] Before entering the loop, the program should create an empty integer slice
of size (length) 3.

During each pass through the loop:
[*] the program prompts the user to enter an integer to be added to the slice.
[*] The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in
sorted order.

[*] The slice must grow in size to accommodate any number of integers which the user decides to enter.

[*] The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	intsLenInit := 3
	ints := make([]int, intsLenInit)
	nints := 0
	fmt.Println(ints)
	scn := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Current ints: ", ints)
		fmt.Println("Gimme an int (X to quit):")
		if scn.Scan() {
			inp := scn.Text()
			if strings.ToLower(inp) == "x" {
				fmt.Println("Done!")
				break
			}
			valid, err := strconv.Atoi(inp)
			if err != nil {
				fmt.Printf("'%s' not a valid integer\n", inp)
			} else {
				if nints <= intsLenInit-1 {
					fmt.Println("insert at ", nints)
					ints[0] = valid
				} else {
					ints = append(ints, valid)
				}
				nints += 1
				slices.Sort(ints)
			}
		}
	}
}
