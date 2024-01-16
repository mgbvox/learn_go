package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//n := runtime.GOMAXPROCS(1) // First test
	//Result：0120120120120120120122012012012012012012012012012012012012012012012...

	n := runtime.GOMAXPROCS(2) // Second test
	//Result：1122012021002200211011011111220011221021012102010122120012002210122...

	fmt.Printf("The number of cores = %d\n", n)
	time.Sleep(time.Second * 2)

	for {
		go fmt.Print(0)
		go fmt.Print(1)
		go fmt.Print(2)

	}

}
