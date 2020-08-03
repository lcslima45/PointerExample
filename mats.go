package main

import (
	"fmt"
)

func gcd(x *int, y *int) {
	for *y != 0 {
		*x, *y = *y, *x%*y
	}

}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	a := 4
	b := 5
	gcd(&a, &b)
	fmt.Println(a, b)
	fmt.Println(fib(123))
}
