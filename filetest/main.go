package main

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func testCreateDir() {

	err = os.Mkdir("tmp", 0755)
	if err != nil {
		fmt.Printf("%v,%T\n", err, err)
		log.Fatalln(err)
	}
	fmt.Println(newFile)
}

func testCreate() {

	newFile, err = os.Create("goTest.txt")
	if err != nil {
		fmt.Printf("%v,%T\n", err, err)
		log.Fatalln(err)
	}
	fmt.Println(newFile)
	err = newFile.Close()
}

func testFileExit() {
	_, err = os.Stat("goTest.txt")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("文件存在")
}

func testRead() {
	file, _ := os.Open("goTest.txt")
	bytes := make([]byte, 10)
	source := make([]byte, 0)
	len1, _ := file.Read(bytes)
	fmt.Printf("%d\n", len1)
	//copy,字节
	for 0 < len1 {
		for i := 0; i < len1; i++ {
			source = append(source, bytes[i])
		}
		len1, _ = file.Read(bytes)
	}
	fmt.Printf("%s\n", source)
	err = file.Close()
}

func testWrite() {
	//打开一个文件
	//os.O_CREATE 表示文件不存在就会创建。
	//os.O_APPEND 表示以追加内容的形式添加。
	//os.O_WRONLY 表示只读模式。
	//os.O_RDWR 表示读写模式。
	//os.O_EXCL 与O_CREATE一起使用，文件不能存在。
	//os.O_SYNC i/o同步的方式打开。
	//os.O_TRUNC 如果可能，请在打开时截断文件。
	newFile, err = os.OpenFile("goTest.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	//写入字节流
	n, err := newFile.Write([]byte("go你好"))
	if err != nil {
		log.Fatalln(err)
	}
	//写入字符串
	m, err := newFile.WriteString("世界")
	if err != nil {
		log.Fatalln(err)
	}
	//在指定的偏移处(offset)写入内容
	_, err = newFile.WriteAt([]byte("！！！"), int64(n+m))
	if err != nil {
		log.Fatalln(err)
	}
	err = newFile.Close()
}

func main() {

	//创建目录
	//testCreateDir()

	//创建文件
	//testCreate()

	//写入文件
	//testWrite()

	//判断文件是否存在
	//testFileExit()

	//读取文件
	testRead()

}
