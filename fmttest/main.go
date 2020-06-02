package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("请输入姓名和年龄，使用空格分隔")
	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
	fmt.Println("second 请输入姓名和年龄，使用空格分隔")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println(name, age)
}
