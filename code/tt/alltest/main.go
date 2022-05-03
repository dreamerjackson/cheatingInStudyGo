package main

import "fmt"

type kk struct {
	m int64
}

func main() {
	return
	defer func() {
		fmt.Println("123")
	}()

}
