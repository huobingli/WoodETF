package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 打开MySQL连接
	db, err := sql.Open("mysql", "golang:123456@tcp(127.0.0.1:3306)/golang?charset=utf8")
	checkError(err)
	fmt.Println("connected")
	// ?为占位符，防止SQL注入
	stmt, err := db.Prepare("insert into user(username, gender, email, ctime, utime) values(?, ?, ?, ?, ?)")
	checkError(err)
	timestamp := time.Now().Unix()
	// 向占位符传参，Exec是一个不定参数函数，传入的参数与Prepare阶段设置的占位符相等
	res, err := stmt.Exec("zhangsan", 1, "zhangsan@qq.com", timestamp, timestamp)
	checkError(err)
	// 获取生成的数据id
	id, err := res.LastInsertId()
	checkError(err)
	fmt.Println(id)

	stmt, err = db.Prepare("update user set email=? where id=?")
	checkError(err)
	res, err = stmt.Exec("zhangsan@163.com", id)
	checkError(err)
	aff, err := res.RowsAffected()
	checkError(err)
	fmt.Println("affected", aff)

	rows, err := db.Query("select id, username, email from user")
	for rows.Next() {
		var id int
		var username string
		var email string
		err = rows.Scan(&id, &username, &email)
		checkError(err)
		fmt.Println(id, username, email)
	}

	stmt, _ = db.Prepare("delete from user where id>?")
	res, _ = stmt.Exec(1)
	aff, _ = res.RowsAffected()
	fmt.Println("affected", aff)
}
