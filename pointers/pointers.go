package main

import "fmt"

//let's learn about pointers! woohoo!

// zeroval func has an int parameter, so arguments will be passed to it by value.
// it gets a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// zeroptr func has an *int parameter, meaning that it takes an int pointer.
// the *iptr code in this function body then dereferences the pointer form its memory address to the current valueof that address
// assigning a value to a dereferenced pointer changes the value at the referenced address!
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// setting initial variable to use in example
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// & symbol gives us the memory address of i, i.e the pointers to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointers can also be printed too!
	fmt.Println("pointer:", &i)

	//zeroval does not change i in main, but zeroptr does because it is a reference to the memory address for that variable!
}
