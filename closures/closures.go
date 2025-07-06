package main

import "fmt"

// go supports anonymous functions, which can form closures.
// anon funcs are useful when you want to define a functions inline

// the func intSeq returns another function -> we defined it anonymously in the body of intSeq
// the returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() // set the value of nextInt to the returned value from intSeq

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // I didn't know that you could create a new function and copty previous existing one to it!
	fmt.Println(newInts())

}
