// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// License: The MIT License

// In Go it’s idiomatic to communicate errors via an explicit,
// separate return value. This contrasts with the exceptions
// used in languages like Java and Ruby and the overloaded
// single result / error value sometimes used in C.
// Go’s approach makes it easy to see which functions return errors
// and to handle them using the same language constructs employed
// for any other, non-error tasks.
package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		return -1, errors.New("can't work with 42")
	}
	// A nil value in the error position indicates that there was no error.
	return arg + 3, nil
}

// It’s possible to use custom types as errors by implementing the Error() method on them.
// Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
type argError struct {
	arg int
	prob string
}

// struct instantiation and have a method named Error()
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// In this case we use &argError syntax to build a new struct,
// supplying values for the two fields arg and prob.
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	fmt.Println(f1(42)) // Can not working
	fmt.Println(f1(43)) // Working
	fmt.Println(f2(42)) // Can not working

	// The two loops below test out each of our error-returning functions.
	// Note that the use of an inline error check on the if line is a common idiom in Go code.
	for _, v := range []int{7, 42} { // for a value of the slice
		if r, e := f1(v); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, v := range []int{7, 42} { // for a value of the slice
		if r, e := f2(v); e != nil {
			fmt.Println("f2 failed")
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error type via type assertion.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
// Visit http://blog.golang.org/2011/07/error-handling-and-go.html for more on error handling.
