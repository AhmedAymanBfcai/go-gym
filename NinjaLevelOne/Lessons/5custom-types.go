package main

import "fmt"

var a int

// Creating custom type:
type hotdob int

var b hotdob

func main() {
	a = 40
	b = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// You can not do a = b [Differnent types]
}
