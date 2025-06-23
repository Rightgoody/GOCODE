package main

import "fmt"

func main() {
	fmt.Println("basic loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("classic condidtion loop like in c++")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("loop x amount of times specified after the range keyword. Output is \"range\" + i")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("infinite loop until breaks. Reminds me of while(true)")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("range loop again but with condition inside and a continue keyword")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
