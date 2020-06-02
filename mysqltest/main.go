package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//数据库连接配置
const (
	username = "root"
	password = "root@123"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

//定义DB数据库连接池
var DB *sql.DB

type Info struct {
	id   int64
	name string
	addr string
}

func initDB() {
	//连接字符串：用户名:密码@tcp(IP:PORT)/数据库?charset=utf8
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println("=====", path)
	DB, _ = sql.Open("mysql", path)
	//最大闲置连接时间
	DB.SetConnMaxLifetime(100)
	//最大空闲连接
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connect success")
}

//插入数据
/*
	1.开启事务
	2.预编译
	3.执行sql语句
	4.提交事务
	5.处理返回结果
*/
func insert(info Info) bool {
	fmt.Println(info)
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Failed to open transaction")
	}
	//预编译
	sql := "insert into info(`name`,`addr`) values (?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//执行sql
	res, err := stmt.Exec(info.name, info.addr)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	id, _ := res.LastInsertId()
	row, _ := res.RowsAffected()
	fmt.Println("======完成====== id=", id)
	return row > 0
}

//查询所有数据
func selectInfo() (infos []Info) {
	sql := "select * from info"
	query, err := DB.Query(sql)
	if err != nil {
		fmt.Println("query error")
		return
	}
	for query.Next() {
		var info Info
		err = query.Scan(&info.id, &info.name, &info.addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		infos = append(infos, info)
	}
	return
}

//查询根据主键查询
func selectInfoById(id int64) (info Info) {
	sql := "select * from info where id = ?"
	DB.QueryRow(sql, id).Scan(&info.id, &info.name, &info.addr)
	return
}

//根据主键删除数据
func deleteInfoById(id int64) (isOk bool) {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Failed to open transaction")
	}
	sql := "delete from info where id = ?"
	stmt, err := tx.Prepare(sql)
	result, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("Exec error：", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Failed to commit transaction", err)
		return
	}
	row, _ := result.RowsAffected()
	isOk = row > 0
	return
}

func main() {
	initDB()

	//info := Info{name: "数据2name", addr: "golang"}
	//insert(info)

	//infos := selectInfo()
	//fmt.Println(infos)

	//id := int64(12)
	//info := selectInfoById(id)
	//fmt.Println(info)

	//id := int64(12)
	//deleteInfoById(id)
}
