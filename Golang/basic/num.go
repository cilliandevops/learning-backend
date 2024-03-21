// @Time    : 2024/3/20 14:37:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 8
	var b int32 = 4
	var c int64 = 4
	d := 4
	var e float32 = 1.2313113222
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	//查看数据类型
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(e)
}