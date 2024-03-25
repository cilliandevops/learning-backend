// @Time    : 2024/3/21 11:12:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

func main() {
	// 声明一个有5个元素的整型数组，所有元素初始化为默认值0
	var myArray [5]int
	fmt.Println("空数组:", myArray)

	// 使用索引为数组的某个位置赋值
	myArray[0] = 100
	myArray[4] = 500
	fmt.Println("赋值数组:", myArray)

	// 直接在声明时初始化数组
	initializedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("初始化数组:", initializedArray)

	// 使用for循环遍历数组
	for i, value := range initializedArray {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	//数组的遍历
	//[...]表示不限长度的数组,但实际还是根据数组元素的数量来定义长度
	var arr3 = [...]string{"北京", "上海", "深圳"}
	//第一种遍历方式
	for i := 0; i < len(arr3); i ++ {
		fmt.Println(arr3[i])
	}
	//第二种遍历方式
	for index, value := range arr3 {
		fmt.Println(index, value)
	}

	// 使用len函数获取数组长度
	fmt.Println("数组长度:", len(myArray))

	//定义时赋值
	var c = [1]string{"Tony"}
	fmt.Println(c)
}

