package main

import "fmt"

func main() {
	fmt.Println("traditional if/else statement")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	fmt.Println("\nsingle if statement")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	fmt.Println("\nlogical && or || operators can be used in the conditionals")
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is/are even")
	}

	fmt.Println("\nstatement can precede conditionals. any variable declared in this statement can be used in current or subsequent branches")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is a single digit number")
	} else {
		fmt.Println(num, "is a multi-digit number")
	}

	fmt.Println("\nthere are no ternary operators in golang!!")
}
