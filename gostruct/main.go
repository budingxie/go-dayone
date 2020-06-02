package main

import "fmt"

type structName struct {
	title  string
	author string
	bookId int
}

func main() {
	type user struct {
		name string
		age  int
	}
	//创建一个结构体
	fmt.Println(structName{"Go 语言", "张三", 6575})

	//key-->>value形式创建
	fmt.Println(structName{title: "go", author: "李四", bookId: 1587})

	fmt.Println(structName{title: "go", bookId: 1587})

	//创建和访问
	var sn1 = structName{"Go 语言", "张三", 6575}
	fmt.Println(sn1.title)
	sn1.title = "语言"
	fmt.Println(sn1.title)

	//结构体指针
	var structPointer *structName
	fmt.Println(structPointer)
	structPointer = &sn1
	fmt.Printf("%x\n", structPointer)
	fmt.Println(*structPointer)

	//结构体指针访问方式
	sn := new(structName)
	sn.title = "dd"
	(*sn).bookId = 123
	fmt.Println(sn)

}
