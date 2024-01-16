package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func subsort(subarray []int, c chan []int, wg *sync.WaitGroup) {
	slices.Sort(subarray)
	// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
	fmt.Print("Sorted subarray: ", subarray, "\n")
	c <- subarray
	wg.Done()
}

func chunkSlice(sli []int, nChunks int) [][]int {
	sliLen := len(sli)
	chunkSize := (sliLen / nChunks) + 1

	chunks := make([][]int, 0)
	for i := 0; i < sliLen; i += chunkSize {
		end := min(i+chunkSize, sliLen)
		sub := sli[i:end]
		chunks = append(chunks, sub)
	}
	return chunks
}

func prompt(args ...string) string {
	var msg string
	quitseq := "ctrl+c"
	terminus := ">"
	switch len(args) {
	case 1:
		msg = args[0]
	case 2:
		msg = args[0]
		quitseq = args[1]
	}

	formatted := strings.TrimLeft(fmt.Sprintf("%s (%s to quit) %s ", msg, quitseq, terminus), " ")
	fmt.Print(formatted)
	scn := bufio.NewScanner(os.Stdin)
	result := ""
	if scn.Scan() {
		inp := scn.Text()
		clean := strings.Trim(strings.ToLower(inp), " ")
		if clean == quitseq && quitseq != "" {
			fmt.Println("Quitting")
			os.Exit(0)
		} else {
			result = clean
		}
	}
	return result
}

/*
This could have just been a merge sort problem, guys :(
*/

func main() {
	//Write a program to sort an array of integers.
	wg := sync.WaitGroup{}

	c := make(chan []int, 2048)

	// The program should prompt the user to input a series of integers.
	inp := prompt("Gimme some ints, separated by whitespace e.g. 1 34 6 1...")
	arr := make([]int, 0)
	for _, elem := range strings.Split(inp, " ") {
		val, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not convert value to int64: %s", inp))
		}
		arr = append(arr, int(val))
	}

	// Or if you don't want to do that, just uncomment this
	//arr := []int{5, 2, 3, 9, 1, 4, 5, 67, 3, 1, 4, 6, 7, 4, 3, 23, 5, 6, 75, 3, 2}

	//The program should partition the array into 4 parts.
	//Each partition should be of approximately equal size.
	for _, chunk := range chunkSlice(arr, 4) {
		wg.Add(1)
		// each slice is sorted by a different goroutine.
		go subsort(chunk, c, &wg)
	}
	// wait for them to complete
	wg.Wait()
	// close the channel so we can iterate over it
	close(c)

	// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
	all := make([]int, 0)
	// Merge them
	for sorted := range c {
		all = append(all, sorted...)
	}
	// Sort them (this is redundant work, but the question doesn't specify how our merge should sort things, so...).
	slices.Sort(all)

	// When sorting is complete, the main goroutine should print the entire sorted list.
	fmt.Println("Sorted ints:")
	fmt.Println(all)
}
