package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length   int
	chartStr string
)

const (
	//纯数字
	chooseNum string = "num"
	//纯英文字母
	chooseLetter string = "letter"
	//特殊字符
	chooseSpecial string = "special"
	//所有
	chooseAdvanced string = "advanced"
	//密码原始串
	numStr = "0123456789"
	str    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	speStr = `!"#$%&'()*+,-"`
)

func testFlag1() string {
	//创建密码源：不同类型的密码，字符串不一样
	var sourceStr string
	flag.IntVar(&length, "L", 6, "设置密码长度")
	flag.StringVar(&chartStr, "T", "num", "选择密码类型")
	//启动解析参数
	flag.Parse()
	fmt.Printf("密码长度=%d;密码类型=%s\n", length, chartStr)
	//判断密码类型
	if chooseNum == chartStr {
		sourceStr = fmt.Sprintf("%s", numStr)
	} else if chooseLetter == chartStr {
		sourceStr = fmt.Sprintf("%s", str)
	} else if chooseSpecial == chartStr {
		sourceStr = fmt.Sprintf("%s", speStr)
	} else if chooseAdvanced == chartStr {
		sourceStr = fmt.Sprintf("%s%s%s", numStr, str, speStr)
	} else {
		sourceStr = numStr
	}
	//随机种子
	rand.Seed(time.Now().UnixNano())
	//创建密码切片
	password := make([]byte, 0)
	for i := 0; i < length; i++ {
		size := rand.Intn(len(sourceStr))
		password = append(password, sourceStr[size])
	}
	pass := string(password)
	return pass
}

func main() {
	password := testFlag1()
	fmt.Println(password)
}
