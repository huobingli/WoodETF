package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	// defer Db.Close() // 注意这行代码要写在上面err判断的下面 ??
}

type Person struct {
	A int `db:"a"`
	B int `db:"b"`
	C int `db:"c"`
	D int `db:"d"`
	E int `db:"e"`
}

func main() {
	var person []Person

	err := Db.Select(&person, "select a, b, c, d, e from test where id=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)

	// stmt, err := db.Prepare("alter table dev convert to character set utf8 collate utf8_general_ci")  //要修改一下编码
	// if stmt != nil {
	//     stmt.Exec()
	//     stmt.Close()
	// }

	// _ = err
}
