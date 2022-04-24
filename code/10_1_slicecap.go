package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 从下标2 一直到下标4，但是不包括下标4
	numbers1 := numbers[2:4]
	fmt.Println("len:", len(numbers1), "cap:", cap(numbers1))
	numbers4 := numbers[2:4:4]
	fmt.Println("len:", len(numbers4), "cap:", cap(numbers4))
}
