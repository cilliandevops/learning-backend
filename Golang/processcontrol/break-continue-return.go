// @Time    : 2024/3/25 16:23:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

func main() {
	//break,跳出整个for循环
	//当满足某个条件，for循环不再需要，则使用break跳出整个for循环
	k := 1
	for {
		if k <= 5 {
			fmt.Print(k)
		} else {
			fmt.Println("跳出循环\n")
			break
		}
		k ++
	}
	//continue,跳出本次循环
	//当i=2的时候，continue直接跳出当前这次循环，进行到i=3的循环
	//当i=4 ...
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}
	//return 跳出当前函数
	fmt.Println("\n -----")

	for i := 0; i < 10; i ++ {
		fmt.Println(i)
		if i == 5 {
			return
		}
	}
	test()
	fmt.Println("最后执行的打印")
}

func test() {
	fmt.Println("这里是test函数1")
	return
	fmt.Println("这里是test函数2")
}