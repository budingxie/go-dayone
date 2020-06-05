package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Email  string `mcl:"email"`
	Name   string `mcl:"name"`
	Age    int    `mcl:"age"`
	Github string `mcl:"github" default:"a8m"`
}

func testReflect1() {
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
		//获取tag的值
		fmt.Println(f.Tag.Get("mcl"), f.Tag.Get("default"))
	}
}

func testReflect2() {
	u := &User{Name: "Ariel Mashraki"}
	v := reflect.ValueOf(u).Elem()
	f := v.FieldByName("Name")
	//确保field 是有效的 也是可以写的
	//if !f.IsValid() || !f.CanSet() {
	//	return
	//}
	//if f.Kind() != reflect.String || f.String() != "" {
	//	return
	//}
	f.SetString("张三")
	fmt.Println(u)
}

func main() {
	//testReflect1()
	testReflect2()
}
