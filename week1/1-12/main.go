/**
* @Author : jiahongming
* @Description :
* @Time : 2021/11/3 1:14 上午
* @File : main
* @Software: GoLand
**/
package main

import "fmt"

func main() {

	//map的三种定义方式
	//第一种
	m := make(map[string]string)
	fmt.Println(m)
	//第二种
	var m2 map[string]string
	fmt.Println(m2)
	//第三种
	m3 := map[string]string{
		"鲁菜": "大肠",
		"川菜": "宫保鸡丁",
	}
	fmt.Println(m3)
	//复合类型map
	m4 := map[string]map[int]string{}
	fmt.Println(m4)

	m5 := map[int]string{}
	m5[1] = "A"
	m5[2] = "B"
	m5[3] = "C"
	m5[4] = "D"
	m5[5] = "E"
	//map的循环
	for idx, item := range m5 {
		fmt.Println(idx, item)
	}
	//map是无序的
	//keys: chiken orange apple

	//获取map的某一个值
	v2, ok := m5[2]
	if ok {
		fmt.Println(v2)
	} else {
		fmt.Println("not find")
	}

	// map delete
	delete(m5, 1)
	fmt.Println(m5)

	//删除不存在元素也不会报错
	delete(m5, 10)
	fmt.Println(m5)
}
