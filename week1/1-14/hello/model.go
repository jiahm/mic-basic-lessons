/**
* @Author : jiahongming
* @Description :
* @Time : 2021/11/3 10:55 上午
* @File : model
* @Software: GoLand
**/
package model

import "fmt"

func Cook(f Food) {

}

func Send(f Food) {

}

type MyInt int32

type Hello func(string)

type Person struct {
	Food Food
}

type Food struct {
	Name string
}

func (p *Person) SayHello(name string) {
	fmt.Println(name + "吃了吗？" + p.Food.Name)
}

//菜系
type FoodSystem struct {
	Name string
}

type Food2 struct {
	FoodName string
	FoodSystem
	Another FoodSystem
}
