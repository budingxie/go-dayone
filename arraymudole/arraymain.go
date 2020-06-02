package main

import (
	"fmt"
)

func main() {
	//申明数组
	var array [10]float32
	var i int
	fmt.Println(array)

	var balance = [5]float32{12.3, 12.4, 12, 12.4, 64.6}
	for i = 0; i < len(balance); i++ {
		fmt.Printf("%f \t", balance[i])
	}
	fmt.Println()
	var balance1 = [...]float32{1, 32.3, 23.1}
	fmt.Println(balance1, len(balance1))

	fmt.Println(balance[0], balance[4])
	fmt.Println(balance1[0], balance1[len(balance1)-1])

}
