package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order,
from least to greatest. Use your favorite search tool to find a description
of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort()
which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements
are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation.

Your Swap() function should take two arguments:
	a slice of integers
	index value i which indicates a position in the slice.

The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

func Swap(arr []int, idx int) {
	if idx < 0 || idx > len(arr)-2 {
		// can't swap before 0 or after the end!
		return
	}
	storage := arr[idx+1]
	arr[idx+1] = arr[idx]
	arr[idx] = storage
}

func BubbleSort(arr []int) {
	// only look at first len(arr) -1 elements,
	// since can't swap past the end of the slice!
	window := arr[:len(arr)-1]
	for {
		didSwap := false // initialize negative
		for idx, elem := range window {
			// we're in a swapable position, and
			// the next value is less than the current value
			if elem > arr[idx+1] {
				// swap 'em
				Swap(arr, idx)
				// note that a swap occurred
				didSwap = true
			}
		}
		// No swaps happened? We're sorted!
		if !didSwap {
			break
		}
	}
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Please provide a list of space separated ints, unsorted.")
	fmt.Println("E.g. '1 43 5898 0 -6 314'")
	fmt.Println("Input:")
	if scn.Scan() {
		inp := scn.Text()
		if inp == "" {
			fmt.Println("Invalid input - got an empty string.")
			return
		}
		var ints []int
		for _, elem := range strings.Split(inp, " ") {
			asInt, err := strconv.Atoi(elem)
			if err != nil {
				fmt.Println("The element you provided is not a valid integer:", elem)
				return
			}
			ints = append(ints, asInt)

		}
		BubbleSort(ints)
		fmt.Println("Here's your ints - we sorted them for you!")
		fmt.Println(ints)
	} else {
		fmt.Println("No input provided...")
	}
}
