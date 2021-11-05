/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 6:35 下午
* @File : main
* @Software: GoLand
**/
package main

import (
	"fmt"
	"net/http"
)

func VisitUrl(url string) (int, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		fmt.Printf("%d", res.StatusCode)
		return res.StatusCode, nil
	}
}

func SwitchShow(url string) {
	if code, err := VisitUrl(url); err != nil {
		fmt.Println(err)
	} else {
		switch code {
		case 200:
			fmt.Println("\n请求成功")
			//fallthrough
		case 404:
			fmt.Println("\n地址不存在")
		default:
			panic("\n未知错误")
		}
	}
}
func main() {
	SwitchShow("https://www.baidu.com")
}
