package main

import "fmt"

type mytype int

var x mytype

var y int

func main() {
	// Exercise #5
	y = int(x) // The conversion

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
