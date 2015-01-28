// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License

// Go supports recursive functions. Here’s a classic factorial example.
package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 { // If n eq 0 then break
		return 1
	}
	return n * fact(n - 1) // This is a loop that Call fact()
}

func main() {
	fmt.Println(fact(17))
}
