package main

import (
	"fmt"
	"strings"
)

func main() {
	console := ConsoleLog{}
	var log Log
	log = &console
	log.Debug("act={}, desc={}", "Debug", "log print")
}

type Level string

type Log interface {
	Init(filePath string, fileName string)
	SetLevel()
	Debug(format string, arg ...string)
	Info(format string, arg ...string)
	Error(format string, arg ...string)
}

//控制台输出
type ConsoleLog struct {
}

func (con *ConsoleLog) Init(filePath string, fileName string) {
}
func (con *ConsoleLog) SetLevel() {
}
func (con *ConsoleLog) Debug(format string, arg ...string) {
	for _, val := range arg {
		format = strings.Replace(format, "{}", val, 1)
	}
	fmt.Println(format)
}
func (con *ConsoleLog) Info(format string, arg ...string) {

}
func (con *ConsoleLog) Error(format string, arg ...string) {

}

//文件输出
type FileLog struct {
}

func (con *FileLog) Init(filePath string, fileName string) {

}
func (con *FileLog) Debug(format string, arg ...string) {

}
func (con *FileLog) Info(format string, arg ...string) {

}
func (con *FileLog) Error(format string, arg ...string) {

}
