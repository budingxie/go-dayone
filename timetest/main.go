package main

import (
	"fmt"
	"time"
)

/*
//Month--实现原理
type name int
const (
	zhangsan name = 123
)
func (m name) String() string {
	return "zhangsan"
}
*/

func testTime1() {
	//获取当前时间，类型是Go的时间类型Time
	currentTime := time.Now()
	//年
	year := currentTime.Year()
	//月
	month := currentTime.Month()
	//日
	day := currentTime.Day()
	//小时
	hour := currentTime.Hour()
	//分钟
	minute := currentTime.Minute()
	//秒
	second := currentTime.Second()
	//纳秒
	nanosecond := currentTime.Nanosecond()

	fmt.Printf("%v,%T\n", currentTime, currentTime)
	fmt.Printf("%v,%T\n", year, year)
	//默认打印字符串
	fmt.Printf("%d,%T\n", month, month)
	fmt.Printf("%v,%T\n", day, day)
	fmt.Printf("%v,%T\n", hour, hour)
	fmt.Printf("%v,%T\n", minute, minute)
	fmt.Printf("%v,%T\n", second, second)
	fmt.Printf("%v,%T\n", nanosecond, nanosecond)
	//根据年月日时分秒,
	currentDate := time.Date(year, month, day, hour, minute, second, nanosecond, time.Local)
	fmt.Printf("%v,%T\n", currentDate, currentDate)
	i := currentDate.Unix()
	fmt.Printf("%v,%T\n", i, i)
	//获取当前时间戳,单位是秒
	unix := currentTime.Unix()
	//获取当前时间戳,单位是纳秒
	nunix := currentTime.UnixNano()
	fmt.Printf("%v,%T\n", unix, unix)
	fmt.Printf("%v,%T\n", nunix, nunix)
	//把当前时间格式化
	format := currentTime.Format("2006-01-02 15:04:05")
	fmt.Printf("%v,%T\n", format, format)
	//已知时间戳，转化为时间
	t := time.Unix(unix, 0)
	fmt.Printf("%v,%T\n", t, t)
	//已知字符串转时间
	formatStr := "2020-05-06 15:22:03"
	parse, err := time.Parse("2006-01-02 15:04:05", formatStr)
	fmt.Printf("%v,%T\n", parse, parse)
	fmt.Printf("%v,%T\n", err, err)
}

func main() {
	testTime1()
}
