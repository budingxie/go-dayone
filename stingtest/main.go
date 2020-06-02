package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	string转[]byte
	//var data []byte = []byte(str)

	[]byte转string
	//var str string = string(data[:])
*/

func stringTest() {
	str := "中国abc123"
	strSize := []rune(str)
	fmt.Printf("str=%s,len(str)=%d,chart=%d\n", str, len(str), len(strSize))
	fmt.Println(utf8.RuneCountInString(str))
}

func stringsTest() {
	str1 := "abcD"
	str2 := "abcd"
	//按照字典序比较大小，排在前面的小，后面的大
	has := strings.Compare(str1, str2)
	fmt.Printf("%v", has)
	fmt.Printf("%v\n", str1 >= str2)
	//忽略大小写
	fold := strings.EqualFold(str1, str2)
	fmt.Printf("%v\n", fold)
	u := str1[0]
	fmt.Printf("%v,%T\n", u, u)

	//判断字符串中是否包含子串
	s := "abcdef"
	sustr := "bcdf"
	contains := strings.Contains(s, sustr)
	fmt.Println(contains)

}
func main() {

	//stringTest()
	//stringsTest()
	str := "123中国"
	//strSize := [] rune(str)
	var runes []int32 = []rune(str)
	fmt.Printf("%T", runes)
	for index, val := range runes {
		fmt.Printf("a[%d]=%v\n", index, val)
	}
}
