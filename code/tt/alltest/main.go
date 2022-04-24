package main

import "fmt"

type kk struct {
	m int64
}

func main() {
	for i := 0; ; i++ {
		a := kk{int64(i)}
		fmt.Println(a)
	}
}
