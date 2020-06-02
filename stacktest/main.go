package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type Stack []interface {
}

//长度
func (stack Stack) Len() int {
	return len(stack)
}

//容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//压栈
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//获取
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

//弹出
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

//中缀表达式转化为后缀表达式
func processExpression(exp string) string {
	stack := Stack{}
	postfix := ""
	expLen := len(exp)
	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		switch char {
		case " ":
			continue
		case "(":
			//左括号入栈
			stack.Push(char)
		case ")":
			//右括号弹出，直到遇到左括号
			for stack.Len() != 0 {
				preChar, _ := stack.Top()
				if preChar == "(" {
					//弹出
					stack.Pop()
					break
				}
				postfix += preChar.(string) + "#"
				stack.Pop()
			}
			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			//j小于总长度，判断字符是否是一个十进制的数字字符
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit + "#"
			//当遇到操作符之后或者越界，最后多执行了一步j++,所有需要去除
			i = j - 1
		default:
			//当前操作符号，遇到高有优先级的运算符号，不断弹出，直到遇到更低优先级运算符
			for stack.Len() != 0 {
				top, _ := stack.Top()
				top1 := top.(string)
				if top1 == "(" || isLower(top1, char) {
					break
				}
				postfix += top1 + "#"
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}
	// 栈不空则全部输出
	for stack.Len() != 0 {
		pop, _ := stack.Pop()
		postfix += pop.(string) + "#"
	}
	return postfix
}

func computerNum(resData []string) (total int) {
	computerStack := Stack{}
	for _, val := range resData {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			com1, _ := computerStack.Pop()
			com2, _ := computerStack.Pop()
			switch val {
			case "+":
				v1, _ := strconv.Atoi(com1.(string))
				v2, _ := strconv.Atoi(com2.(string))
				total = v2 + v1
				computerStack.Push(strconv.Itoa(total))
			case "-":
				v1, _ := strconv.Atoi(com1.(string))
				v2, _ := strconv.Atoi(com2.(string))
				total = v2 - v1
				computerStack.Push(strconv.Itoa(total))
			case "*":
				v1, _ := strconv.Atoi(com1.(string))
				v2, _ := strconv.Atoi(com2.(string))
				total = v2 * v1
				computerStack.Push(strconv.Itoa(total))
			case "/":
				v1, _ := strconv.Atoi(com1.(string))
				v2, _ := strconv.Atoi(com2.(string))
				total = v2 / v1
				computerStack.Push(strconv.Itoa(total))
			}
		} else {
			computerStack.Push(val)
		}
	}
	return
}

//比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

//对数值开根号计算
func Sqrt(x float64) float64 {
	z := 1.0
	for {
		tmp := z - (z*z-x)/(2*z)
		if tmp == z || math.Abs(tmp-z) < 0.000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func main() {
	sqrt := Sqrt(0.25)
	fmt.Println(sqrt)
	//fmt.Println("请输入计算表达式，输入完成之后回车就能获取到结果，只能计算整数")
	//var computerStr string
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("表达式有错误，请重新输入")
	//		for {
	//			fmt.Scanln(&computerStr)
	//			expression := processExpression(computerStr)
	//			resData := strings.Split(expression, "#")
	//			total := computerNum(resData)
	//			fmt.Printf("%s=%d", computerStr, total)
	//			fmt.Println()
	//		}
	//	}
	//}()
	//for {
	//	fmt.Scanln(&computerStr)
	//	expression := processExpression(computerStr)
	//	resData := strings.Split(expression, "#")
	//	total := computerNum(resData)
	//	fmt.Printf("%s=%d", computerStr, total)
	//	fmt.Println()
	//}

}
