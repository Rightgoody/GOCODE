package main

import "fmt"

// factorial function
func fact(n int) int {
	//base case
	if n == 0 {
		return 1
	}
	//recurse case
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// fibonacci function
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
