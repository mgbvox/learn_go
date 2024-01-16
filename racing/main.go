package main

import "fmt"

func add(sli *[]float64) {
	end := len(*sli)
	fib := (*sli)[end-2] + (*sli)[end-1]
	*sli = append(*sli, fib)
}

/*
This is a faulty implementation of the Fibonacci algorithm which suffers from race conditions.
*/

func main() {

	// We create a float slice here, which will be our shared state (and means of communication)
	sli := []float64{0, 1, 1}

	N := 30
	for i := 0; i < N; i++ {
		/*
			This is where problems start.

			We pass a *pointer* to the slice to our add function, which modifies it in place.
			We spawn N go routines that will execute this in an arbitrary order.

			The add() func gets the last two elements of the slice, adds them together, and
			appends them to the list. Were this to be executed synchronously, we'd get the first
			N (really N+3) Fibonacci numbers.

			You can test this by switching which line below is commented out.

			In the concurrent case, the len(sli) calculation in add() is not guaranteed to reflect the
			*actual* length of sli when we get to the append() step; as such, we pretty frequently
			wind up performing redundant operations earlier in the slice than where its head actually
			is.

			Run the concurrent version several times (you may need to make N larger depending on your
			computer hardware; 30 is a good value for my machine). The final printed list will vary in
			size!
		*/
		go add(&sli)
		//add(&sli)
	}

	fmt.Println(sli)

}
