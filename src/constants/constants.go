// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Go supports constants of character, string, boolean, and numeric values.
package main

import "fmt"
import "math"
// Or using:
// import (
// 	"fmt"
// 	"math"
// )

// const declares a constant value
const s string = "constants" // string is a type


func main() {
	fmt.Println(s) // Return constants

	// A const statement can appear anywhere a var statement can
	const n = 50000000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit cast
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// For example, here math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
