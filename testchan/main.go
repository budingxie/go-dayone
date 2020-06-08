package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 1
	i := <-c
	fmt.Println(i)
}
