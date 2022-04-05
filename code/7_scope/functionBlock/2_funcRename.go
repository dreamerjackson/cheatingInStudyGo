package main

import "fmt"

//1\就近原则
var x int = 5

func main() {

	var x int = 99
	x = 100
	fmt.Println("testx", x) //100

}
