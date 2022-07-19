// Fill in the blanks so this program compiles and produces
// the output shown.
package main

import "fmt"

func main() {
	// Declare a variable that holds a map with strings for keys
	// and booleans for values.
	var isVowel ____
	// Call make to create the map.
	isVowel = ____
	// Assign true to the key "a", and false to the keys "b" and "c".
	____ = true
	____ = false
	____ = false
	fmt.Printf("%#v\n", isVowel) // => map[string]bool{"a":true, "b":false, "c":false}


	// Create a map without make with ints as keys and strings as values. The 1 key
	// should have the value "H", 2 should have the value "He", and
	// 3 should have the value "Li".
	var elements ____
	elements = 
	// Print all the keys and corresponding values in the slice.
	// Order doesn't matter.
	for ____ := range elements {
		fmt.Println(atomicNumber, symbol)
	}
	// => 3 Li
	// => 1 H
	// => 2 He
}
