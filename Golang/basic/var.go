// @Time    : 2024/3/19 15:52:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"
var age = 20
//全局定义不限制位置，但建议统一放在函数代码前面
func main() {
	//基本声明
	//var 变量名 类型 = 值
	//声明并初始化
	var name string = "cilliandevops"
	var age1 int = 18
	var China bool = true
	fmt.Println(name, age1, China)


	//类型推断方式定义变量，可以忽略类型
	var age2 = 18
	var name1 = "cilliandevops"

	//局部变量，只能用在函数体内，同一个代码块内不能多次声明同一个变量
	age3 := "18"
	fmt.Printf("%T %T\n", age,age1,age2,age3,name1)

	fmt.Println(age)
	//一次定义多个变量
	var username, sex string
	username = "cillian"
	sex = "man"
	fmt.Println(username, sex)
	//
	var a, b = 1, 2
	fmt.Println(a,b)
	//常用方式，定义多个变量
	var (
		user string
		city = "chengdu"
	)
	fmt.Println(user,city)
}




