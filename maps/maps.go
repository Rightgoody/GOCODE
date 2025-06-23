package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int) // create a map [key-type]value-type

	//creating elements within the map
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) //prints entire map contents

	//obtaining value from key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if key doesnt exist, zero value of type is instead used.I guess that means key gets created and initialized to zero?
	v3 := m["k3"]
	fmt.Println("k3:", v3)

	fmt.Println("len:", len(m)) // obtaining the number of elements in map

	// deleting a key/value pair from the map
	delete(m, "k2")
	fmt.Println("map:", m)

	// deleting all key/value pairs from the map
	clear(m)
	fmt.Println("map:", m)

	/*
		The optional second return value when getting a value from a map indicates
		if the key was present in the map. This can be used to disambiguate between
		missing keys and keys with zero values like 0 or "".
		Here we didn’t need the value itself, so we ignored it with the blank identifier _
	*/
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	//declaring and initializing the map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//usuing "equal" utility function
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
