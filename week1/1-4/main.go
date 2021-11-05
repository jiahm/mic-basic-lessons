/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 2:52 下午
* @File : main.go
* @Software: GoLand
**/
package main

import (
	"fmt"
)

func main() {
	var price int32 = -68 //
	fmt.Println(price)
	var price2 uint32 = 68 //uint不能是负数
	fmt.Println(price2)

	var price3 float32 = 99.99 //float32 float64
	fmt.Println(price3)

	var checked bool = false
	fmt.Println(checked)

	var name string = "从0到Go语言微服务架构师路线图"
	fmt.Println(name)

	//byte
	fmt.Println([]byte(name))

	//rune对标 char
	//complex()
	name4, price4, num4 := "炒菜", 68, 1
	var total float64 = 0
	discount := 0.7
	// 显示类型转换
	total = float64(price4) * float64(num4) * discount
	fmt.Printf("%s总价%f\n", name4, total)

}
