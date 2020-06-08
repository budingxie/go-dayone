package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Conf struct {
	Mysql MysqlConf `ini:"mysql"`
	Redis RedisConf `ini:"redis"`
}

type MysqlConf struct {
	Password string `ini:"password"`
	Username string `ini:"username"`
	Url      string `ini:"url"`
}

type RedisConf struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

func main() {
	data, err := ioutil.ReadFile("./conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	conf := &Conf{}
	err = UnmarshalConf(data, conf)
	if err != nil {
		fmt.Printf("解析错误，err：%v", err)
	}
	fmt.Printf("get conf: %#v", conf)
}

func UnmarshalConf(data []byte, conf *Conf) (err error) {

	lineStr := strings.SplitN(string(data), "\r\n", -1)
	var structName string
	confField := reflect.TypeOf(conf).Elem()
	confVal := reflect.ValueOf(conf).Elem()
	for _, line := range lineStr {
		//1.验证数据是否合法,是否需要解析
		hasLegal, checkErr := checkLine(line)
		//1.1非法数据
		if checkErr != nil {
			err = checkErr
			return
		}
		//1.2不需要解析的数据，如注释
		if !hasLegal {
			continue
		}
		//2.解析结构体对应的字段
		if len(line) > 2 && line[0] == '[' && line[len(line)-1] == ']' {
			line = strings.TrimSpace(string(line[1 : len(line)-1]))
			for i := 0; i < confField.NumField(); i++ {
				tagName := confField.Field(i).Tag.Get("ini")
				if tagName == line {
					structName = confField.Field(i).Name
				}
			}
			fmt.Printf("line: %s, structName: %s\n", line, structName)
			continue
		}
		//3.获取结构体下的字段属性
		fieldLien := strings.SplitN(line, "=", 2)
		fmt.Println(fieldLien)
		structField := confVal.FieldByName(structName)
		elem := structField.Type()
		var elemName string
		for j := 0; j < elem.NumField(); j++ {
			tagName := elem.Field(j).Tag.Get("ini")
			if tagName == fieldLien[0] {
				elemName = elem.Field(j).Name
			}
		}
		//4.赋值操作
		elemVal := confVal.FieldByName(structName).FieldByName(elemName)
		switch elemVal.Kind() {
		case reflect.String:
			elemVal.SetString(fieldLien[1])
		case reflect.Int:
			parseInt, _ := strconv.ParseInt(fieldLien[1], 10, 64)
			elemVal.SetInt(parseInt)

		}
	}
	return
}

func checkLine(line string) (hasLegal bool, err error) {
	hasLegal = true
	return
}
