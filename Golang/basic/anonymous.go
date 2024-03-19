// @Time    : 2024/3/19 17:10:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!

// 匿名变量
package main
import (
"fmt"
)
//GetClass 获取班级信息
func GetClass() (stuNum int, className, headTeacher string) {
	return 29, "一班", "cillian"
}
func getFullName() (string, string) {
	return "Cillian", "Devops"
}
func main () {
	stuNum, _, _ := GetClass()
	//只想获取班级人数setNum ，其他不需要
	fmt.Println(stuNum)
	// 只使用第一个返回值，忽略第二个返回值
	firstName, _ := getFullName()
	fmt.Println("名字:", firstName)

	// 只使用第二个返回值，忽略第一个返回值
	_, lastName := getFullName()
	fmt.Println("姓氏:", lastName)
}