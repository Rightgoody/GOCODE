package main

import (
	"fmt"
	"math"
)

// declaring variables outside of function, like in asm .data section(?)

const s string = "constant" // explicitly declaring its type

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
