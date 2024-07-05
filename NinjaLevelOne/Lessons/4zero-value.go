package main

import "fmt"

var y string = "Fares"
var a int

func main() {
	// Zero Values:
	// 0 for for int, 0.0 for float, "" for string
	// nil for (pointers, functions, interfaces, slices, channels , maps)
	fmt.Println("Zero value video:")
	fmt.Println(y)        // Fares
	fmt.Printf("%T\n", y) // string

	fmt.Println(a)        // 0
	fmt.Printf("%T\n", a) // int
}
