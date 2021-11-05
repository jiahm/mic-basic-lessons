/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 5:07 下午
* @File : main
* @Software: GoLand
**/
package main

import "fmt"

const (
	//枚举
	//iota 默认值为0，每行累加1
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

	Author = "jhm"
	Book   = "Go语言极简一本通"
)

func main() {
	//常量只读不允许修改
	const version = 1.0
	const appName = "面向加薪学习"

	fmt.Printf("%s-%s-%f-%s", Author, Book, version, appName)
	fmt.Println(Friday)
}
