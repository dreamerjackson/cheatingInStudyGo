package main

import (
	"fmt"
	"time"
)

func main() {
	// 完整的 C 风格的for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------")

	// 只有条件判断的for循环
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
	fmt.Println("----------")

	// for-range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	fmt.Println("----------")

	// 无限循环
	for {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
	}
}
