package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:111111@tcp(47.114.171.118:32334)/ark_test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {
	init()

	// var person []Person
	// err := Db.Select(&person, "select * from arkk_etf", 1)
	// if err != nil {
	// 	fmt.Println("exec failed, ", err)
	// 	return
	// }

	// fmt.Println("select succ:", person)
}
