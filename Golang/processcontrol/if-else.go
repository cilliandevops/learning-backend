// @Time    : 2024/3/25 11:48:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

func main() {
	//第一种,变量在if条件判断之外,变量作用域在整个main函数中
	score := 99
	if score >= 90 {
		fmt.Println("优秀")
	} else if (score > 75) {
		fmt.Println("良好")
	} else {
		fmt.Println("及格")
	}

	//第二种，变量在if条件判断中，变量作用域在if函数中
	if score := 99; score >= 90 {
		fmt.Println("优秀")
	} else if (score > 75) {
		fmt.Println("良好")
	} else {
		fmt.Println("及格")
	}


}

