package main

import (
	"fmt"
	"time"
)

func main() {

	const a string = "abc"
	const b = "abc"
	time.Sleep(time.Second * 5)
	fmt.Println(a, b)
}
