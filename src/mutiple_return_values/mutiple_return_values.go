// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.
package main

import (
	"fmt"
)

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Here we use the 2 different return values from
	// the call with multiple assignment.
	a, b := vals()
	fmt.Println("a is", a)
	fmt.Println("b is", b)

	// If you only want a subset of the returned values, use the blank identifier _
	_, c := vals()
	fmt.Println("c is", c)
	d, _ := vals()
	fmt.Println("d is", d)
}

