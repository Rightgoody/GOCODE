package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Print("Construct a struct of type person: ")
	fmt.Println(person{"Bob", 20}, "\n") // construct a struct of type person

	fmt.Print("Can name the fields when initializing struct: ")
	fmt.Println(person{name: "Alice", age: 30}, "\n") // can name the fields when initilaizing struct

	fmt.Print("Fields omitted will be zero-valued: ")
	fmt.Println(person{name: "Fred"}, "\n") // Fields omitted will be zero-valued

	fmt.Print("Prefixing struct with & will create a pointer to struct object: ")
	fmt.Println(&person{name: "Ann", age: 40}, "\n") //prefixing the struct type with & will make a pointer to the struct

	fmt.Print("It's idiomatic to encapsulate new struct creation in constructor functinos: ")
	fmt.Println(newPerson("Jon"), "\n") // It's idiomatic to encapsulate new struct creation in constructor functions

	// Interacting with structs:
	fmt.Println("\nInteracting with Structs: \n")

	fmt.Print("Accessing struct fields with dot operator: ")
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, "\n") // accessing struct fields with dot operator

	fmt.Print("Can also create pointers to structs and access fields using said dot operator- referenced object will automatically be dereferenced: ")
	sp := &s
	fmt.Println(sp.age, "\n") // can also create pointers to struct and access fields using dot operator

	fmt.Print("Structs are mutable: ")
	sp.age = 51
	fmt.Println(sp.age, "\n")

	//From example: If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
	fmt.Print("Single value struct types do not have to be named. Value can be anon struct type- known as \"table-driven tests\": ")
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
