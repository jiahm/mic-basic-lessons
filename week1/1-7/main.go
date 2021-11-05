/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 7:21 下午
* @File : main
* @Software: GoLand
**/
package main

import (
	"fmt"
	"time"
)

//go语言没有while
func main() {
	sum := 0
	// 第一种写法
	for i := 0; i < 100; i++ {
		sum += i
	}

	//第二种写法
	i := 0
	for ; i < 100; i++ {
		sum += i
	}

	//第三种写法
	for ; ; i++ {
		if i == 100 {
			break
		}
		sum += i
	}

	//第四种写法
	for {
		if i == 100 {
			break
		}
		i++
		sum += i
	}

	fmt.Printf("%d\n", sum)

	for {
		fmt.Println("从0到Go语言微服务架构师", i)
		i++
		time.Sleep(time.Duration(1) * time.Second)
	}
}
