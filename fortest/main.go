package main

import (
	"fmt"
	"sort"
	"time"
)

func test(a []int) {
	for index, value := range a {
		fmt.Printf("a[%d]=%v\n", index, value)
		//睡眠3秒
		time.Sleep(3 * time.Second)
	}
}

func testTest() {
	a := [3]int{5, 2, 3}
	b := a[:]
	sort.Ints(b)
	test(b)
}

func testCustomSort(a []int) {
	/*
		i：0，	1，		2		...		n
		j：n，	n-1		n-2		...		0
		y=n+(n-1)+(n-2)+...+3+2+1
		y=1+2+3+...+(n-2)+(n-1)+n
		2*y=(n+1)+(n+1)+...+(n+1)
		y=(n+1)*n/2
	*/

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				var temp int = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
}
func testCustomSortTest() {
	a := [...]int{5, 2, 3, 6, 3, 4, 55, 67, 83, 2, 9}
	fmt.Println(a)
	testCustomSort(a[:])
	fmt.Println(a)
}

func main() {
	//testTest()
	testCustomSortTest()

}
