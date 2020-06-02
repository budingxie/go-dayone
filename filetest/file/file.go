package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件输入输出
func testWriteFile() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
		"1中文测试\n",
	}
	file, err := os.Create("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for _, p := range proverbs {
		//file类型实现了io.Writer
		n, err := file.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")
}
func testReadFile() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	p := make([]byte, 4)

	var chunks []byte
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		if 0 == n {
			break
		}
		fmt.Print(string(p[:n]))
		chunks = append(chunks, p[:n]...)
	}
	fmt.Print(string(chunks[:len(chunks)]))
}

//标准输入输出
func testStdout() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _, p := range proverbs {
		// 因为 os.Stdout 也实现了 io.Writer
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
}
func testStdin() {
	var buffer []byte = make([]byte, 4)
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(n)
	fmt.Println(string(buffer[:]))
}

//io拷贝
func testIoCopy() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("channels orchestrate mutexes serialize \n")
	file, err := os.Create("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// io.Copy完成了从proverbs读取数据并写入file的流程
	_, err = io.Copy(file, proverbs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file created")
}
func testIoStdout() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
func testIoWriteString() {
	file, err := os.Create("./magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = io.WriteString(file, "Go is fun!")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

//管道copy
func testPipWrite() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	pipeR, pipeW := io.Pipe()
	go func() {
		defer pipeW.Close()
		io.Copy(pipeW, proverbs)
	}()
	//从另一端pipeR中读取
	io.Copy(os.Stdout, pipeR)
	pipeR.Close()
}

func testBufIo() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//按行读取
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Print(line)
	}

}

func testIoUtil() {
	bytes, err := ioutil.ReadFile("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", bytes)
	err = ioutil.WriteFile("./planetsCopy.txt", bytes, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	//testWriteFile()
	//testReadFile()
	//testStdout()
	//testStdin()
	//testIoCopy()
	//testIoStdout()
	//testIoWriteString()
	//testPipWrite()
	//testBufIo()
	//testIoUtil()
}
