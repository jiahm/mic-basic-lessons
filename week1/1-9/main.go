/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/31 2:49 下午
* @File : main
* @Software: GoLand
**/
package main

import "fmt"

func byValue(price int) {
	price += 20
}

func byRef(price *int) {
	*price += 20
}
func main() {
	var price int = 68
	var ptr *int = &price
	fmt.Println(ptr)

	*ptr = 88
	fmt.Println(price)

	byValue(price)
	fmt.Println(price)

	byRef(&price)
	fmt.Println(price)

}
