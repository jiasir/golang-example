// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Go has various value types including
// strings, integers, floats, booleans, etc. Here are a few basic examples.
package main

import "fmt"

func main() {
	// Go has various value types including strings, integers, floats, booleans, etc.
	// Here are a few basic examples.

	// Strings, which can be added together with +
	fmt.Println("go" + "lang") //Return golang

	// Integers and floats
	fmt.Println("1+1 =", 1+1) //Return 1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) //Return 7.0/3.0 = 2.3333333333333335

	// Booleans, with boolean operators as you’d expect
	fmt.Println(true && false) // Return false
	fmt.Println(true || false) // Return true
	fmt.Println(!true) // Return false
}

