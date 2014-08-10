// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Our first program will print the classic “hello world” message. Here’s the full source code.

// Define a package name main, Every Go program is made up of packages.
package main

// This program is using the packages with import "fmt".
import "fmt"

// Package start running in package main.
func main() {
	fmt.Println("hello world") // Return hello world
}

// To run the program, put the code in hello-world.go and use go run.
// $ go run hello-world.go

// Sometimes we’ll want to build our programs into binaries. We can do this using go build.
// $ go build hello-world.go
// $ ls
// hello_world	hello_world.go

// We can then execute the built binary directly.
// $ ./hello-world
// hello world

// Now that we can run and build basic Go programs.
