// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License


// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a
// function inline without having to name it.
package main

import "fmt"

// This function intSeq returns another function,
// which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to
// form a closure.
func intSeq(num0 int) func(num1 int) int { // func function_name(arg) anonymous_function(arg) return_value {
	i := num0 // i is a counter
	return func(num1 int) int {
		i += num1 // Counter + num1
		return i
	}
}

// This is a no sequence function
func noSeq() int { // func function_name(arg) return_value {
	i := 0 // counter
	i += 1 // counter+1
	return i // Return 1
}

func main() {
	// Initialize intSeq assign to nextInt variable
	nextInt := intSeq(0) // 0 is num0
	fmt.Println(nextInt(1)) // 1 is num1, the output is 1
	fmt.Println(nextInt(1)) // The output is 2

	// Initialize intSeq assign to nextInt variable
	newInt := intSeq(0) // 0 is num0
	fmt.Println(newInt(1)) // 1 is num1, the output is 1
	fmt.Println(newInt(1)) // The output is 2

	noClosure := noSeq() // assign noSeq() to noClosure variable
	fmt.Println(noClosure) // The output is 1
	fmt.Println(noClosure) // The output is 1
}
