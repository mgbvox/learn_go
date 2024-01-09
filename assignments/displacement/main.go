package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func prompt(msg string) float64 {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println(msg)
	if scn.Scan() {
		inp := scn.Text()
		val, err := strconv.ParseFloat(inp, 64)
		if err != nil {
			panic(fmt.Sprintf("Could not convert value to float64: %s", inp))
		}
		return val
	} else {
		panic("No input!")
	}
}

func GenDisplacementFn(a, v0, x0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((.5 * a) * math.Pow(t, 2)) + (v0 * t) + x0
	}
}

func main() {
	fn := GenDisplacementFn(prompt("acceleration?"), prompt("initial velocity?"), prompt("initial displacement?"))
	t := prompt("time (sec)?")
	fmt.Printf("Displacement after time %f: %f\n", t, fn(t))
}
