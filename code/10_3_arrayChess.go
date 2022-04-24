package main

import "fmt"

func main() {
	var arr = [5][5]int{}
	arr[2][2] = 1
	arr[2][3] = 2
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			switch arr[i][j] {
			case 0:
				fmt.Printf("%3s", ".")
			case 1:
				fmt.Printf("%3s", "X")
			case 2:
				fmt.Printf("%3s", "O")
			}
		}
		fmt.Println()
	}
}
