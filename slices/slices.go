package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "meow"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("length:", len(s))

	s = append(s, "d")
	var test []string = s
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("TESTapd:", test)

	//copy slices
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sli:", l)

	l = s[:5]
	fmt.Println("sl1:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
