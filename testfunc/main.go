package main

import "fmt"

func main() {
	var ff method = fm
	ff(1)
}

type method func(int)

func fm(i int) {
	fmt.Println(i)
}
