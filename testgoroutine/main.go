package main

import (
	"fmt"
	"time"
)

func runWork(c <-chan int) {
	for {
		data := <-c
		fmt.Println(data)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	c := make(chan int, 1)
	go runWork(c)
	go func(ct chan<- int) {
		for i := 1; ; i++ {
			ct <- i
		}
	}(c)
	time.Sleep(10 * time.Second)
}
