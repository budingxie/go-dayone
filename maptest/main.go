package main

import (
	"fmt"
)

func mapTest1() {
	var a map[string]string = make(map[string]string, 16)
	fmt.Println(a)
	a["key1"] = "value1"
	s, ok := a["key1"]
	if ok {
		fmt.Println("the value of key is ", s)
		a["key1"] = "value2"
	} else {
		fmt.Println("the value of key is not exit")
	}
	fmt.Println(a)
}

func mapTest2() {

	a := make(map[string]string)

	fmt.Printf("the map is %v  length %d %p\n", a, len(a), a)
	a["key1"] = "value1"
	a["key2"] = "value2"
	a["key3"] = "value3"
	a["key4"] = "value4"
	a["key5"] = "value5"
	a["key6"] = "value6"
	a["key7"] = "value7"
	a["key8"] = "value8"
	a["key9"] = "value9"
	fmt.Println(a)
	for key, value := range a {
		if key == "key2" {
			delete(a, key)
		}
		fmt.Printf("%s=%s\n", key, value)
	}
	fmt.Println(a)
}

func main() {

	//mapTest1()
	mapTest2()

}
