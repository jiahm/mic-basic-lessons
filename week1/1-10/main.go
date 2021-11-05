/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/31 3:19 下午
* @File : main
* @Software: GoLand
**/
package main

import "fmt"

func printFood(foods [3]string) {
	foods[2] = "大饼"
	fmt.Println(foods)
}
func main() {
	//数组的几种定义方式
	var array1 [6]string
	array2 := [3]string{"火锅", "烧烤", "家常菜"}
	array3 := [...]string{"火锅", "烧烤", "家常菜"} //编译器处理有多大

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	var matrix [3][4]string
	var matrix2 [3][4]int
	fmt.Println(matrix)  //[[   ] [   ] [   ]]
	fmt.Println(matrix2) //[[0 0 0 0] [0 0 0 0] [0 0 0 0]]

	//数组第一种遍历方式
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	//数组第二种遍历方式
	for idx, item := range array2 {
		fmt.Println(idx, item)
	}

	//for (类型 变量：集合){} java
	//for 变量 in 集合  python
	printFood(array2)
	fmt.Println(array2)
}
