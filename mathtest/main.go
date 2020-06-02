package main

import (
	"fmt"
	"math"
)

func main() {
	i := 100
	//取绝对值
	abs := math.Abs(float64(i)) //i2 := int(abs)
	//向上取整,5.0是5; 5.1是6
	ceil := math.Ceil(5.1)
	//线下取整
	floor := math.Floor(5.9)
	//取余数，和取模一样
	mod := math.Mod(11, 6)
	//取整数和取小数
	modf, frac := math.Modf(5.26)
	//x的y次方
	pow := math.Pow(2, 3)
	//10的n次方
	pow10 := math.Pow10(3)
	//开平方
	sqrt := math.Sqrt(16)
	//开立方
	cbrt := math.Cbrt(8)
	//Pi
	pi := math.Pi
	fmt.Printf("%v, %T\n", abs, abs)
	fmt.Printf("%v, %T\n", ceil, ceil)
	fmt.Printf("%v, %T\n", floor, floor)
	fmt.Printf("%v, %T\n", mod, mod)
	fmt.Printf("%v, %T\n", modf, modf)
	fmt.Printf("%v, %T\n", frac, frac)
	fmt.Printf("%v, %T\n", pow, pow)
	fmt.Printf("%v, %T\n", pow10, pow10)
	fmt.Printf("%v, %T\n", sqrt, sqrt)
	fmt.Printf("%v, %T\n", cbrt, cbrt)
	fmt.Printf("%v, %T\n", pi, pi)
}
