package main

import "fmt"

// Use var for : zero values, package scope
// DECLARE there is a VARIABLE with the IDENTIFER "Z" and give it TYPE int
// ASSIGNS the zero value of TYPE int to "z"
var z int // Output is 0

func main() {
	// The var keyword

	var c = 100    // The diffs between var and short delcartion are that you can delcare a variable outside the func body by using var keyword and the scope you can use variables in also
	fmt.Println(c) // Output is 100
	// fmt.Println(z) // Output is 0

	fmt.Println(z) // Output is 0 (the zero value)
}
