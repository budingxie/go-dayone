package main

import (
	"fmt"
)

//栈结构
type stack struct {
	value []int32
	len   int
}

//压栈
func (st *stack) push(val int32) {
	if st.len == len(st.value) {
		st.value = append(st.value, val)
		st.len = st.len + 1
		return
	}
	st.value[st.len] = val
	st.len = st.len + 1
}

//出栈
func (st *stack) pull() (isHave bool, val int32) {
	length := st.len
	if length == 0 {
		return false, 0
	}
	st.len = length - 1
	val = st.value[st.len]
	return true, val
}

var (
	signStack *stack
	numStack  *stack
)

//1.中缀表达式转化为后缀表达式
func processExpression(fistExpression string) (lastExpression string) {
	//得到字符
	dataByte := []rune(fistExpression)
	for _, val := range dataByte {
		//判断该字符是什么
		if val == '+' || val == '-' || val == '*' || val == '/' || val == '(' {
			fmt.Print("#")
			isHave, sign := signStack.pull()
			if !isHave {
				signStack.push(val)
				continue
			}
			if val == '(' {
				continue
			}
			//比较sign和val的优先级
			//1.val优先级高，直接入栈；
			//2.val优先级等于或者低，弹出sign，遇到1.条件，或者栈为空停止
			if compLevel(val, sign) {
				//sign优先级高
				for {
					fmt.Print(sign)
					isHave, sign = signStack.pull()
					if !isHave {
						signStack.push(val)
						continue
					}
					if compLevel(val, sign) {
						break
					}
				}
			} else {
				signStack.push(val)
			}
		} else if val == ')' {
			for {
				isOk, sign := signStack.pull()
				if !isOk {
					fmt.Println("error ！！！")
					break
				}
				if sign == '(' {
					break
				}
				fmt.Print(string(sign))
			}
		} else {
			fmt.Print(string(val))
		}
	}
	for {
		break
	}
	return lastExpression
}

func compLevel(val int32, sign int32) (ok bool) {
	//比较sign和val的优先级
	//1.val优先级高，直接入栈；
	//2.val优先级等于或者低，弹出sign，为空停止
	if sign == '*' || sign == '/' {
		//sign优先级高
		return true
	} else {
		if val == '*' || val == '/' {
			//sign优先级低
			return false
		}
		//sign优先级高
		return true
	}
}

func main() {

	computerStr := "1*2+15-(3-1)*2"
	//1#2#15#3#1#-#2*-+*
	signStack = &stack{
		value: make([]int32, 0),
		len:   0,
	}
	processExpression(computerStr)
	//fmt.Println(string(numStack.value[:numStack.pointer]))
	//fmt.Println(string(signStack.value[:signStack.pointer]))

}
