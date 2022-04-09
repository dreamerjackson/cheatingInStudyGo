package main

import "fmt"

func main() {
	fmt.Println(((1 + 2) * 5) / 3) // 5
	fmt.Println((1 + 2*5) / 3)     // 3

	fmt.Println(2 * 2 & 2)        //  4 & 2  100 & 010 -> 0
	fmt.Println(2 * (2 & 2) << 1) // 8
}
