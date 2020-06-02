package main

import (
	"fmt"
)

var str string = "123"

func main() {
	/*
		var age int;
		var b bool = true;
	*/
	var (
		age int
		bo  bool = true

		//无符号
		a uint8
		b uint16
		c uint32
		d uint64

		//有符号
		e int8
		f int16
		g int32
		h int64

		//浮点型
		f1 float32
		f2 float64

		//字符类型
		str string
	)
	var vName string
	vName, vName1 := "123s", 12
	fmt.Println(age, bo)
	fmt.Printf("%s: a=%d, b=%d, c=%d, d=%d\n", "无符号", a, b, c, d)
	fmt.Printf("%s: e=%d, f=%d, g=%d, h=%d\n", "无符号", e, f, g, h)
	fmt.Printf("%s: f1=%f, f2=%f\n", "浮点型", f1, f2)
	fmt.Println(str, "=======")
	fmt.Println(vName, vName1)
}
