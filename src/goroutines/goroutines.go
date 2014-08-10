// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// A goroutine is a lightweight thread of execution.
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s).
	// Here’s how we’d call that in the usual way, running it synchronously.
	f("fucking direct")

	// Using goroutine
	go f("fucking goroutine")

	// Start goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("anonymous") // () arg for calling

	// Our two goroutines are running asynchronously in separate goroutines now,
	// so execution falls through to here.
	// This Scanln code requires we press a key before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
