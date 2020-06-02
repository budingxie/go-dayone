package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//testStringsReader()
	customStringsReader()
}

func testStringsReader() {
	var reader *strings.Reader = strings.NewReader("clear is better than clever ")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

func customStringsReader() {
	read := newAlphaReader("custom strings read hello")
	p := make([]byte, 4)
	for {
		n, err := read.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

type alphaReader struct {
	//资源
	src string
	//当前读取到的位置
	cur int
}

//创建一个实例
func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

//过滤函数
func alpha(r byte) byte {
	if (r > 'A' && r < 'Z') || (r > 'a' || r < 'z') {
		return r
	}
	return 0
}

//Read方法
func (a *alphaReader) Read(p []byte) (int, error) {
	//当前位置 >= 字符串长度  说明已经读取到结尾 返回EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	//x表示剩下未读取得长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		//剩下长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		//剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		// 每次读取一个字节，执行过滤函数
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}
