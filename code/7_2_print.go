package main

import "fmt"

// Hello world
func main() {
	a := 10
	b := "zjx"
	c := 3.1415926
	d := true
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // string
	fmt.Printf("%T\n", c) // float64
	fmt.Printf("%T\n", d) // bool
}
