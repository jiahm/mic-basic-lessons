/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 1:47 下午
* @File : main.go
* @Software: GoLand
**/
package main

import "fmt"

//作用域在包内部，并不是全局的作用域，go没有全局作用域的变量。
var (
	book    = "\n Go语言极简一本通"
	lesson1 = "Go语言微服务架构核心22讲"
	lesson2 = "从0到Go语言微服务架构师路线图"
)

func main() {
	//变量声明方法一：声明一个变量并初始化
	var price int32 = 68
	var name string = "cai"
	fmt.Println("hello world")
	fmt.Printf("%d,%s", price, name)

	//变量声明方法二：声明多个变量，省去类型并初始化
	var price2, weight = 68, 1
	fmt.Printf("\n%d,%d", price2, weight)

	//变量声明方法三：省去var关键字，通过“:=”进行定义及赋值；
	price3 := 68
	weight2 := 1
	fmt.Printf("\n%d,%d", price3, weight2)

	fmt.Printf("%s", book)
	fmt.Println(lesson1)
	fmt.Println(lesson2)
}
