/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"



//1\就近原则
var x int=5

func main(){

	var x int = 99;
	x = 100;
	fmt.Println("testx",x) //100

}