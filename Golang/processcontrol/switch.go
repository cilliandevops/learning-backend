// @Time    : 2024/3/25 15:02:00
// @Author  : Cillian
// @Email   : cilliandevops@gmail.com
// Website  : www.cillian.website
// Have a good day!
package main

import "fmt"

func main() {
	grade := "B"
	marks := 90

	switch	marks {

	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 60, 70:
		grade = "C"
	default:
		grade ="D"

	}

	//switch  {
	//
	//case marks >= 90:
	//	grade ="A"
	//case marks >= 80:
	//	grade ="B"
	//case marks >= 70 :
	//	grade ="C"
	//case marks >= 60 :
	//	grade ="D"
	//default:
	//	grade = "E"
	//
	//}

	fmt.Printf("你的成绩为%s\n",grade)

	//switch  {
	//case grade == "A":
	//	fmt.Println("成绩优秀")
	//case grade == "B":
	//	fmt.Println("良好")
	//case grade == "C":
	//	fmt.Println("加油")
	//}
}