// @Time    : 2024/3/20 11:02:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

var s = "cilliandevops"
func main() {
	s1 := "hello" +
		"world"
	var s2 = `
		第一行
		第二行
		第三行`
	fmt.Println("字符串是：",s)
	fmt.Println("字符串长度是：",len(s))
	fmt.Println(s1)
	fmt.Println(s2)
}
