package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	var (
		b string  = "字符串"
		c int     = 123
		d float64 = 15.9
	)
	value := reflect.ValueOf(a)
	fmt.Println(value)
	a = b
	fmt.Println(a)
	a = c
	fmt.Println(a)
	a = d
	fmt.Println(a)
}
