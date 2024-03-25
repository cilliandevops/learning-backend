// @Time    : 2024/3/25 14:39:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

func main() {
	//第一种
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//第二种（while 形式）
	i := 0
	for (i < 10) {
		fmt.Println(i)
		i ++
	}
	//无限循环，for
	//for {
	//	fmt.Println(1)
	//}

	//for range 键值对循环
	str := "abc上海"
	for index, val := range str {
		fmt.Printf("索引=%d, 值=%c \n", index, val)
	}

	for i,_ :=range str	{
		fmt.Println(i)
	}

	//switch case
	score := "D"
	switch score {
	case "A":
		fmt.Println("非常棒")
	case "B":
		fmt.Println("优秀")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

