package main

import "fmt"

//在花括号中声明的变量只在花括号中有效。
func main() {
	fmt.Println("Hello World") // x is out of scope
	{                          // x is out of scope
		x := 5         // x is in scope
		fmt.Println(x) // x is in scope
	} // x is out of scope again
}

/*
下面代码无效：

func main() {
    {
        x := 5
    }
    fmt.Println(x)
}
*/
