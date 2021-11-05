/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/30 7:56 下午
* @File : main
* @Software: GoLand
**/
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

//func 函数名称（参数列表） 返回值列表
func ShowName() string {
	return "从0到Go语言微服务架构师"
}

//函数可以返回单个值，返回值匿名。
func ShowInfo(bookName, author string) string {
	return bookName + "-" + author
}

//函数可以返回单个或者多个值，返回值匿名。
func ShowInfoAndPrice(bookName, author string, price float64) (string, float64) {
	return bookName + "-" + author, price
}

//函数可以返回单个或者多个值，可以给返回值取名。
func ShowInfoAndPrice2(bookName, author string, price float64) (bookinfo string, finalprice float64) {
	bookinfo = bookName + "-" + author
	finalprice = price
	return bookinfo, finalprice
}

func VisitUrl(url string) (int, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		//fmt.Printf("%d", res.StatusCode)
		return res.StatusCode, nil
	}
}

func SwitchShow2(url string) (int, error) {
	if code, err := VisitUrl(url); err != nil {
		return 0, err
	} else {
		return code, nil
	}
}

//函数式编程
//do就当作参数
func PrintBookInfo(do func(string, string, float64) (bookInfo string, finalPrice float64), bookName, author string, price float64) {
	bookInfo, finalPrice := do(bookName, author, price)
	fmt.Println(bookInfo)
	fmt.Println(finalPrice)

}

func PrintBookInfo2(do func(string, string, float64) (bookInfo string, finalPrice float64), bookName, author string, price float64) {
	pointer := reflect.ValueOf(do).Pointer()
	methodName := runtime.FuncForPC(pointer)
	fmt.Println(methodName)
	bookInfo, finalPrice := do(bookName, author, price)
	fmt.Println(bookInfo)
	fmt.Println(finalPrice)

}

//匿名函数 do
func PrintBookInfo3(bookName, author string, price float64) {

	do := func(string, string, float64) (bookInfo string, finalPrice float64) {
		bookInfo = bookName + "" + author
		finalPrice = price
		return
	}
	do(bookName, author, price)
}

//可变参数列表作为参数
func ShowAll(foodList ...string) string {
	r := ""
	for _, itme := range foodList {
		r += itme
	}
	return r
}
func main() {
	fmt.Println(ShowName())
	fmt.Println(ShowInfo("从0到Go语言微服务架构师", "jhm"))
	fmt.Println(SwitchShow2("http://www.baidu.com"))

	PrintBookInfo(ShowInfoAndPrice2, "从0到Go语言微服务架构师", "jhm", 99)

	fmt.Println(ShowAll("A", "B", "C"))
}
