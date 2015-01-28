// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License

// Switch statements express conditionals across many branches.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Here’s a basic switch
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same case statement.
	// We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's a before noon")
	default:
		fmt.Println("it's a after noon")
	}
}

