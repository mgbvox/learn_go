package main

import (
	"fmt"
	"math"
)

func GenDispalceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return ((a*math.Pow(t, 2))/2 + v0*t + s0)
	}
	return fn
}

func main() {
	fmt.Println("A time displacement calculation program")
	fmt.Println("Input the acceleration a, intial velocity v0 and intial displacement s0")
	var a, v0, s0 float64
	fmt.Printf("a: ")
	fmt.Scan(&a)
	fmt.Printf("\n")
	fmt.Printf("v0: ")
	fmt.Scan(&v0)
	fmt.Printf("\n")
	fmt.Printf("s0: ")
	fmt.Scan(&s0)
	fmt.Printf("\n")

	sn := GenDispalceFn(a, v0, s0)

	var time float64
	fmt.Printf("Enter a time (in seconds): ")
	fmt.Scan(&time)
	fmt.Printf("\nThe displacement after  %f seconds is %f", time, sn(time))
}

// fn := (a * math.Pow(t,2))/2 + v0 * t + s0
