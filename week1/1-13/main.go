/**
* @Author : jiahongming
* @Description :
* @Time : 2021/11/3 1:40 上午
* @File : main.go
* @Software: GoLand
**/
package main

import "fmt"

type Food struct {
	No   int32
	Name string
}

//go语言没有构造函数，通过工厂函数实现
func NewFood(No int32, Name string) *Food {
	return &Food{
		No:   No,
		Name: Name,
	}
}

//定义一个方法
//f是接收者
func (f Food) show() {
	fmt.Println(f)
}

//定义一个函数
func show(f Food) {
	fmt.Println(f)
}

//通过指针可以改变内容，通过值传递是拷贝的
func (f *Food) SetName(name string) {
	f.Name = name
}
func main() {
	//初始化第一种方式
	var food Food
	food = Food{
		No:   1,
		Name: "chongshaohaishen",
	}
	fmt.Println(food)
	fmt.Println(food.No)
	fmt.Println(food.Name)
	//初始化第二种方式
	food2 := Food{
		No:   2,
		Name: "gongbaojiding",
	}
	fmt.Println(food2)

	food3 := NewFood(3, "回锅肉")
	fmt.Println(food3)
	food3.SetName("宫保鸡丁")
	fmt.Println(food3)
}
