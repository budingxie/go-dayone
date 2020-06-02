package main

import (
	"fmt"
)

func max(num1 int, num2 int) int {

	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	var (
		a = 100
		b = 200
		c int
	)
	c = max(a, b)
	fmt.Println(c)
}
