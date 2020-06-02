package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func testArray() {
	//创建一个数组,其元素的初始值为类型的默认值（0），并且该变量可以直接使用，长度不一样的数组，不能称为一类
	nums := [3]int{}
	nums[0] = 1
	n := nums[0]
	n = 2
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("n: %d\n", n)
}

func testSlice() {
	nums := [3]int{}
	nums[0] = 1
	dNums := nums[:]
	fmt.Printf("dNums: %v, len: %d, cap: %d", dNums, len(dNums), cap(dNums))
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type User struct {
	Email  string `mcl:"email"`
	Name   string `mcl:"name"`
	Age    string `mcl:"age"`
	Github string `mcl:"github" default:"a8m"`
}

func main() {
	//testArray()
	//testSlice()
	var u interface{} = User{}
	//TypeOf 返回 reflect.Type表示u的动态type
	t := reflect.TypeOf(u)
	//kind返回类型
	if t.Kind() != reflect.Struct {
		return
	}
	//变量field
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag.Get("mcl"), f.Tag.Get("default"))
	}
}
