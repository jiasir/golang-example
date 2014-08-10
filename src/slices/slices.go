// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Slices are a key data type in Go,
// giving a more powerful interface to sequences than arrays.
package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("s slice is", s)
	fmt.Println("s len is", len(s))
	fmt.Println("s cap is", cap(s))

	// We can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// In addition to these basic operations,
	// slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append", s)


	// Slices can also be copy’d.
	// Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cps", c)


	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("s[2:5} is", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("s[:5] is", l)

	// And this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("s[2:] is", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	fmt.Println("twoD is", twoD)
	for i := 0; i < 3; i ++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j ++ {
			twoD[i][j] = i + j
		}
	}
	// Note that while slices are different types than arrays,
	// they are rendered similarly by fmt.Println.
	fmt.Println("2d:", twoD)
}

// Visit http://blog.golang.org/2011/01/go-slices-usage-and-internals.html
// for more details on the design and implementation of slices in Go.