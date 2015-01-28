// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.
package main

import "fmt"

// This person struct type has name and age fields
type person struct {
	name string
	age int
}

func main() {
	// To creates a new struct
	fmt.Println(person{"Taio", 26})

	// You can name the fields when initializing a struct
	fmt.Println(person{name: "Taio", age: 26})

	// Omitted fields will be zero-valued
	fmt.Println(person{name:"Taio"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Justin", age: 40})

	// Access struct fields with a dot
	s := person{name: "Rihanna", age: 50}
	fmt.Println(s.name, s.age)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.name, sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
	s.age = 52
	fmt.Println(s.age)
}
