package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

type ARK_ETF struct {
	Ark_Date         string `json:"ark_date" db:"ark_date"`
	Ark_Stock_Name   string `json:"ark_stock_name" db:"ark_stock_name"`
	Ark_Shares       string `json:"ark_shares" db:"ark_shares"`
	Ark_Market_Value string `json:"ark_market_value" db:"ark_market_value"`
	Ark_Weight       string `json:"ark_weight" db:"ark_weight"`
}

type ARK_ETF_STOCK_SHARE struct {
	Ark_Date         string `json:"ark_date" db:"ark_date"`
	Ark_Stock_Name   string `json:"ark_stock_name" db:"ark_stock_name"`
	Ark_Shares       string `json:"ark_shares" db:"ark_shares"`
	Ark_Market_Value string `json:"ark_market_value" db:"ark_market_value"`
}

func get_data(cond string) []ARK_ETF {
	query := fmt.Sprintf("SELECT ark_date,ark_stock_name,ark_shares,ark_market_value,ark_weight FROM %s", cond)
	// fmt.Print(query)
	rows, err := Db.Query(query)
	if err != nil {
		fmt.Printf(`%T`, rows)
		log.Fatal(err)
	}
	defer rows.Close()

	ret := make([]ARK_ETF, 0)
	// fmt.Print(rows)
	for rows.Next() {
		var ark_stock ARK_ETF
		if err := rows.Scan(&ark_stock.Ark_Date, &ark_stock.Ark_Stock_Name, &ark_stock.Ark_Shares, &ark_stock.Ark_Market_Value, &ark_stock.Ark_Weight); err != nil {
			log.Fatal(err)
		}
		// fmt.Print(reflect.Type(rows))
		// fmt.Print(rows)
		fmt.Print(rows)
		ret = append(ret, ark_stock)
	}

	return ret
}

func get_data_count(cond string) []ARK_ETF_STOCK_SHARE {
	query := fmt.Sprintf("SELECT ark_date,ark_stock_name,ark_shares,ark_market_value FROM %s", cond)
	// fmt.Print(query)
	rows, err := Db.Query(query)
	if err != nil {
		fmt.Printf(`%T`, rows)
		log.Fatal(err)
	}
	defer rows.Close()

	ret := make([]ARK_ETF_STOCK_SHARE, 0)
	for rows.Next() {
		var ark_stock ARK_ETF_STOCK_SHARE
		if err := rows.Scan(&ark_stock.Ark_Date, &ark_stock.Ark_Stock_Name, &ark_stock.Ark_Shares, &ark_stock.Ark_Market_Value); err != nil {
			log.Fatal(err)
		}

		ret = append(ret, ark_stock)
	}

	return ret
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
