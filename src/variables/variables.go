// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// In Go, variables are explicitly
// declared and used by the compiler to
// e.g. check type-correctness of function calls.
package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a string = "initial"
	fmt.Println(a) // Return initial

	// You can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c) // Return 1 2

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d) // Return true

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e) // Return 0

	// The := syntax is shorthand for declaring and initializing a
	// variable, e.g. for var f string = "short" in this case.
	f := "shot"
	fmt.Println(f) // Return shot
}

