// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License

// for is Go’s only looping construct. Here are three basic types of for loops.
package main

import (
	"fmt"
)

func main() {
	// The most basic type, with a single condition
	i := 1 // Initial counter
	for i <=3 { // When counter <= 3
		fmt.Println(i) // Print current value
		i = i + 1 // Counter
	}

	// A classic initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// for without a condition will loop repeatedly
	// until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break //break the loop
	}
}


