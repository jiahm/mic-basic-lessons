/**
* @Author : jiahongming
* @Description :
* @Time : 2021/10/31 4:30 下午
* @File : main
* @Software: GoLand
**/
package main

import (
	"fmt"
)

func SoldOut(foods []string) {
	foods[1] = "已售磬"
	fmt.Println(foods)
}
func main() {
	array5 := [...]string{"A", "B", "C", "D", "E", "F", "G"}
	//Slice
	//第一种切片方式
	slice1 := array5[1:3]
	fmt.Println(slice1) //[B C]

	slice2 := array5[1:]
	fmt.Println(slice2) //[B C D E F G]

	slice3 := array5[:3]
	fmt.Println(slice3) //[A B C]

	slice4 := array5[:]
	fmt.Println(slice4) //[A B C D E F G]

	SoldOut(slice4)
	fmt.Println(slice4)

	slice5 := slice4[2:5]
	fmt.Println(slice5)

	slice6 := slice5[2:4]
	fmt.Println(slice6)

	//slice7 := slice5[2:8]
	//fmt.Println(slice7) //panic: runtime error: slice bounds out of range [:8] with capacity 5

	//slice 每次扩容是当前2倍
	slice8 := []string{}
	fmt.Println(len(slice8), cap(slice8))
	for i := 0; i < 10; i++ {
		slice8 = append(slice8, fmt.Sprintf("蛋炒饭%d", i))
		fmt.Println(len(slice8), cap(slice8))
	}

	//遍历切片
	for i, v := range slice8 {
		fmt.Println(i, v)
	}
	//复制切片
	//copy(slice9, slice8)
	//第二种切片方式
	s1 := make([]string, 8)
	fmt.Println(s1)
	//第三种切片方式
	s2 := make([]string, 8, 16)
	fmt.Println(s2)

	s := make([]int, 8)
	//增
	s = append(s, 1)
	//删
	remove(s, 2)
	//改
	s[1] = 1
	//查
	fmt.Println(s[1])

	ss := []string{"炭烤生蚝", "麻辣小龙虾", "干锅鸭"}
	ss2 := make([]*string, len(ss))
	for i, _ := range ss {
		ss2[i] = &ss[i]
	}
	fmt.Println(ss2)
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
