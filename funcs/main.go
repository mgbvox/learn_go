package main

import "fmt"

func foo(y int) {
	y += 1
	fmt.Println("foo y is", y)
}

func boo(y *int) {
	*y += 1
	fmt.Println("boo y is", y)
	fmt.Println("boo *y is", *y)
}

func slifoo(sli []int) []int {
	sli[0] += 1
	return sli
}

func main() {

	x := 2
	foo(x)
	fmt.Println(x)
	boo(&x)
	fmt.Println(x)

	sli := make([]int, 3)
	fmt.Println(slifoo(sli))

}
