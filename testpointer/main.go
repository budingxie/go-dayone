package main

import (
	"fmt"
	"unsafe"
)

func testPointer1() {
	num := 5
	unmPointer := &num
	flNum := (*float32)(unsafe.Pointer(unmPointer))
	fmt.Println(flNum)
}

func testPointer2() {
	type Num struct {
		i string
		j int64
	}
	n := Num{i: "SSADREEEEA", j: 1}

	nPointer := unsafe.Pointer(&n)
	//niPointer := (*string)(unsafe.Pointer(nPointer))
	niPointer := (*string)(nPointer)
	*niPointer = "煎鱼"

	fmt.Println(unsafe.Offsetof(n.j))
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2
	fmt.Printf("n.i: %s, n.j: %d\n", n.i, n.j)

	fmt.Println(unsafe.Offsetof(n.i))
	ni2Pointer := (*string)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.i)))
	*ni2Pointer = "水煮"
	fmt.Printf("n.i: %s, n.j: %d\n", n.i, n.j)

}

func testPointer3() {
	type Num struct {
		i string
		j int64
	}
	n := Num{i: "SSADREEEEA", j: 1}

	njPointer := unsafe.Pointer(&(n.j))

	jPointer := (*int64)(njPointer)
	*jPointer = 153
	fmt.Printf("n.i: %s, n.j: %d\n", n.i, n.j)

	fmt.Println(unsafe.Offsetof(n.j))
	iPointer := (*string)(unsafe.Pointer(uintptr(njPointer) - unsafe.Offsetof(n.j)))
	*iPointer = "1233"
	fmt.Printf("n.i: %s, n.j: %d\n", n.i, n.j)
}

func main() {
	//testPointer1()
	//testPointer2()
	testPointer3()
}
