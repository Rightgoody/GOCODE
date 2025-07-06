package main

import "fmt"

func vals() (int, string) {

	return 3, "meow"
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
