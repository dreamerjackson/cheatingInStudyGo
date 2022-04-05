package main

import "fmt"

func f() {
	fmt.Println(x)
}
func main() {
	f()
}

/*
	$ go build .
	$ ./2_outarg
*/
