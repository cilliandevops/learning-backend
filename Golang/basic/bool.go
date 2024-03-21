// @Time    : 2024/3/20 10:56:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import (
"fmt"
"os"
)
//USER 用户名称
var USER = os.Getenv("USER")
func main() {
	var b bool
	b = (1 != 0)
	//编译时推断变量值
	u := ("cillian" == USER)
	//运行时根据环境判断变量的位
	fmt.Println(u, b)
}
	//返回true