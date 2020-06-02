package main

import "fmt"

func main() {
	var a int = 10
	var pointer *int = &a
	var b int = *&a
	fmt.Printf("变量a的地址=%x\n", &a)
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(b)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	fmt.Println(ptr)

	var ptr1 **int
	fmt.Printf("ptr1 的值为 : %x\n", ptr)
	fmt.Println(ptr1)

	ptr1 = &pointer
	fmt.Printf("ptr1 的值为 : %x\n", ptr)
	fmt.Println(ptr1)
	fmt.Println(**ptr1)

}
