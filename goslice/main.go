package main

import "fmt"

func main() {

	//完整的申明
	var slice1 []int = make([]int, 5)
	fmt.Println(slice1)

	//简写
	slice2 := make([]int, 2)
	fmt.Println(slice2)
	//初始化
	slice2[0] = 2
	slice2[1] = 7

	//追加
	slice2 = append(slice2, 5)
	fmt.Println(slice2)

	//固定长度数组
	s1 := [10]int{1}
	//分片，可变数组
	s2 := []int{1}
	fmt.Println(s1)
	fmt.Println(s2)

	//在后面添加一个元素
	s2 = append(s2, 2)
	fmt.Println(s2)

	//copy
	s3 := make([]int, len(s2), cap(s2)*2)
	copy(s3, s2)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	a := []int{1, 2, 3, 4, 5}
	var i int = 1
	var N int = 2
	fmt.Println(a)
	// 删除中间1个元素
	//a = append(a[:i], a[i+1:]...)
	//fmt.Println(a)
	// 删除中间N个元素
	a = append(a[:i], a[i+N:]...)
	fmt.Println(a)

	// 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+1:])]
	//fmt.Println(a)
	// 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+N:])]
	//fmt.Println(a)
}
