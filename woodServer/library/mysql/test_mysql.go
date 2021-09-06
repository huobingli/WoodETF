package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type ARK_ETF struct {
	Ark_Date         string `db:"ark_date"`
	Ark_Stock_Name   string `db:"ark_stock_name"`
	Ark_Shares       string `db:"ark_shares"`
	Ark_Market_Value string `db:"ark_market_value"`
	Ark_Weight       string `db:"ark_weight"`
}

func get_data(cond string) {
	rows, err := Db.Query("SELECT ark_date,ark_stock_name,ark_shares,ark_market_value,ark_weight FROM ?", cond)
	if err != nil {
		fmt.Printf(`%T`, rows)
		log.Fatal(err)
	}
	defer rows.Close()
	// fmt.Print(rows)
	for rows.Next() {
		var ark_stock ARK_ETF
		if err := rows.Scan(&ark_stock.Ark_Date, &ark_stock.Ark_Stock_Name, &ark_stock.Ark_Shares, &ark_stock.Ark_Market_Value, &ark_stock.Ark_Weight); err != nil {
			log.Fatal(err)
		}
		fmt.Print(ark_stock)
	}
}

// tels
func test_tsla_get_data() {
	rows, err := Db.Query("SELECT ark_date,ark_stock_name,ark_shares,ark_market_value,ark_weight FROM ARKK_ETF WHERE ark_stock_name=?", "TSLA")
	if err != nil {
		fmt.Printf(`%T`, rows)
		log.Fatal(err)
	}
	defer rows.Close()
	// fmt.Print(rows)
	for rows.Next() {
		var ark_stock ARK_ETF
		if err := rows.Scan(&ark_stock.Ark_Date, &ark_stock.Ark_Stock_Name, &ark_stock.Ark_Shares, &ark_stock.Ark_Market_Value, &ark_stock.Ark_Weight); err != nil {
			log.Fatal(err)
		}
		fmt.Print(ark_stock)
	}
}

func init() {

	database, err := sqlx.Open("mysql", "root:111111@tcp(47.114.171.118:32333)/ark_test")
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
	// var person []Person

	// err := Db.Select(&person, "select a, b, c, d, e from test where id=?", 1)
	// if err != nil {
	// 	fmt.Println("exec failed, ", err)
	// 	return
	// }

	// fmt.Println("select succ:", person)

	get_data()
}
