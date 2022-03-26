/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
//函数体内部的变量是function block，注意前后顺序，同时不能跨函数使用。

func main() {
	fmt.Println("Hello World")
	x := 5
	fmt.Println(x)
}

/*
//下面的代码无效：
func main() {
	fmt.Println("Hello World")

	fmt.Println(x)
	x := 5
}


//下面的代码无效2：
func main() {
	x := 5
}
func test(){
    fmt.Println(x)
}
*/