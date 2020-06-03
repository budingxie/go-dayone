package main

import (
	"errors"
	"fmt"
	"io"
)

//type errorString string
//
//func (e errorString) Error() string {
//	return string(e)
//}
//
//func New(text string) error {
//	return errorString(text)
//}
//
//var ErrNamedType = New("EOF")
//var ErrStructType = errors.New("EOF")
//
//func main() {
//	if ErrNamedType == New("EOF") {
//		fmt.Println("Named Type Error")
//	}
//	if ErrStructType == io.EOF {
//		fmt.Println("io Struct Type Error")
//	}
//	if ErrStructType == errors.New("EOF") {
//		fmt.Println("Struct Type Error")
//	}
//}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorString{text}
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		fmt.Println("Error:", ErrType)
	}

	if ErrType == errors.New("EOF") {
		fmt.Println("Errors :", ErrType)
	}

	if io.EOF == io.EOF {
		fmt.Println("io Errors :", ErrType)
	}
	a := 10
	b := 10
	var p1 *int = &a
	var p2 *int = &b
	if p1 == p2 {
		fmt.Println("=====p======")
	}
	if p1 == &a {
		fmt.Println("====a=======")
	}

}
