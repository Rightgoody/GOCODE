package main

import "fmt"

// parameters a and b defined as ints (seperately)
func plus(a int, b int) int {

	return a + b
}

// parameters a b c defined as ints (together)
func plusPlus(a, b, c int) int {

	return a + b + c
}

func main() {

	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
