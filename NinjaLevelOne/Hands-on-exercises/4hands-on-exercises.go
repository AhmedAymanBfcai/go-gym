package main

import "fmt"

type mytype int

var x mytype

func main() {
	// Exercise #4
	fmt.Println(x)        // Output 0 (the zero value)
	fmt.Printf("%T\n", x) // we can use fprint

	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
