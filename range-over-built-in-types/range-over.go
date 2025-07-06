package main

import "fmt"

/*
range iterates over elements in built-in data structures. lets learn some!
*/

func main() {

	// creating a range loop over a int slice
	nums := []int{2, 3, 4}
	sum := 0
	// we use _ underline as placeholder, since we don't need to know the index of the element!
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on array and slices provides both index and value for each entry
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range on map iterates over key/value pairs

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	// the first value is the starting byte index of the rune and the second the rune itself!
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
