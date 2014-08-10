// Name: golang-example
// Author: jiasir (Taio Jia) <jiasir@icloud.com>
// Source: https://github.com/nofdev/go-by-example
// License: The MIT License

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
package main

import "fmt"

func main() {
	// To create an empty map, use the builtin make:
	// make(map[key-type]val-type)
	m := make(map[string]int)
	fmt.Println("m map is", m)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. Println will show all of its key/value pairs
	fmt.Println("m is", m)

	// Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1 is", v1)
	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("m len is", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("delete k2:", m)

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	_, value := m["k2"]
	fmt.Println("value k2", value)

	_, valueK1 := m["k1"]
	fmt.Println("value k1", valueK1)
	keyK1, _ := m["k1"]
	fmt.Println("k1 key is", keyK1)

	// You can also declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	fmt.Println("number", n)
}

