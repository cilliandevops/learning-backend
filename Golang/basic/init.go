// @Time    : 2024/3/20 10:50:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

var inita = "cilliandevops"

func main() {
	fmt.Println("这里是main函数")
}

func init() {
	fmt.Println("这里是init函数")
	fmt.Println(inita)
}



