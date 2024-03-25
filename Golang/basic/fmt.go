// @Time    : 2024/3/21 14:15:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!

package main

import (
"fmt"
)

func main() {
	// 打印输出
	fmt.Println("Hello, World!") // 自动添加换行
	name := "Go"
	version := 1.22
	// 使用Printf进行格式化输出
	fmt.Printf("Language: %s, Version: %.2f\n", name, version)

	// 格式化字符串但不输出到标准输出
	message := fmt.Sprintf("Formatted string: %s", name)
	fmt.Println(message)

	fmt.Printf("String: %s\n", "Hello, Go!")
	fmt.Printf("Integer (decimal): %d\n", 123)
	fmt.Printf("Integer (binary): %b\n", 123)
	fmt.Printf("Integer (hex): %x\n", 456)
	fmt.Printf("Float: %f\n", 78.9)
	fmt.Printf("Float (2 decimals): %.2f\n", 78.9)
	fmt.Printf("Scientific: %e\n", 123400000.0)
	fmt.Printf("Type: %T\n", true)
	//fmt.Printf("Pointer: %p\n", &main)
	fmt.Printf("Quoted string: %q\n", "Hello, Go!")
	fmt.Printf("Go syntax representation: %#v\n", []int{1, 2, 3})
	fmt.Print("zhangsan","lisi","wangwu")
	fmt.Println("zhangsan","lisi","wangwu")
	name1 := "zhangsan1"
	age := 20
	//格式化输出
	fmt.Printf("%s,%d", name1, age)
	//格式化输出，并返回字符串
	info := fmt.Sprintf("%s,%d", name1, age)
	fmt.Println(info)


	// 读取输入
	var input string
	fmt.Println("Enter your name:")
	fmt.Scanln(&input) // 读取输入存储到input变量中
	fmt.Printf("Welcome, %s!\n", input)

//---------------格式化占位符
//`
//	%v: 默认格式的值。对于不同类型的变量，它能智能地选择合适的格式。
//	%#v: 值的Go语法表示。对于变量，它会展示相应的值是如何在Go代码中表示的。
//	%T: 值的类型。
//	%d: 十进制整数。
//	%b: 二进制整数。
//	%o: 八进制整数。
//	%x, %X: 十六进制整数。使用%x时字母为小写，使用%X时字母为大写。
//	%f, %F: 浮点数。%f和%F基本相同，差别在于对于NaN和Infinity的表示。
//	%.2f: 浮点数，保留两位小数。
//	%e, %E: 科学计数法表示的浮点数。%e表示为小写（如2.34e+08），%E表示为大写（如2.34E+08）。
//	%s: 字符串。
//	%q: 双引号围绕的字符串，由Go语法安全地转义。
//	%p: 指针的十六进制表示，前缀为0x。
//
//`


}
