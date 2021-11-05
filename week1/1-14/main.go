/**
* @Author : jiahongming
* @Description :
* @Time : 2021/11/3 10:53 上午
* @File : main
* @Software: GoLand
**/
package main

import (
	"fmt"
	"mic-basic-lessons/week1/1-14/hello"
)

func main() {
	//var foodlist []string // 私有
	var Food model.Food // 公开
	fmt.Println(Food)

	price := 68
	newPrice := model.MyInt(price)
	fmt.Println(newPrice)

	food := model.Food{Name: "豆浆"}

	p := model.Person{Food: food}
	p.SayHello("张三")

	food2 := model.Food2{FoodName: "宫保鸡丁",
		FoodSystem: model.FoodSystem{Name: "鲁菜"},
		Another:    model.FoodSystem{Name: "川菜"},
	}

	fmt.Println(food2.Another.Name)
	fmt.Println(food2.Name)
}
