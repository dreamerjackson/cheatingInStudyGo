package main

import "fmt"

func main() {
	//	定义局部变量：年、月、日
	year := 2008
	month := 12
	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		//判断闰年
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d年%d月的天数为：%d", year, month, days)
}
