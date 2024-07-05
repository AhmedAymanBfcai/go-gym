package main

import "fmt"

var a int

// Creating custom type:
type hotdob int

var b hotdob

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Conversion not casting
	a = int(b) // It will take convert the value assigned to hot dog and convert it to type (a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// You can not do a = b [Differnent types]
}
