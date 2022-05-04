/**
* @Author : jiahongming
* @Description :
* @Time : 2021/12/19 5:49 下午
* @File : main
* @Software: GoLand
**/
package main

import (
	"errors"
	"fmt"
)

func funcRecover() error {
	defer func() { // 匿名函数
		if v := recover(); v != nil {
			fmt.Printf("\n panic recover! v:%v", v)
		}
	}()
	return funCook()
}

func funCook() error {
	panic("停水")
	return errors.New("发生错误了")
}

// 错误处理
func main() {
	// panic recover defer
	err := funcRecover()
	if err != nil {
		fmt.Printf("\n err is %v", err)
	} else {
		fmt.Printf("\n err is nil")
	}
}
