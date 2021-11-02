package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//启动
func main() {
	router := gin.Default()

	// 分组
	v1 := router.Group("/v1")
	{
		v1.GET("/GetStockData/:db/:stock", GetStockData)
		v1.GET("/GetETFData/:db/:time", GetETFData)

		// 获取某etf成分股变化 8.30起
		v1.GET("/GetETFStockChange/:db/:stock", GetETFStockChange)

		// 获取某etf新进成分股数据
		v1.GET("GetETFNewImportStock/:db/:time", GetETFNewImport)
		v1.GET("GetETFNewExportStock/", GetETFNewExport)

		// 获取某etf 所有成分股变化
		v1.GET("GetETFAllStockChange/:db/", GetETFAllStockChange)

		v1.POST("/GetStockData", func(c *gin.Context) {

			stock := c.Query("stock")

			// 获取StockData
			fmt.Printf("stock: %s;", stock)
			c.String(http.StatusOK, "stock: %s;", stock)
		})

		// v1.POST("/GetETFData/:db/:time", GetETFData)
		v1.POST("/GetETFData", func(c *gin.Context) {

			etf := c.Query("etf")

			// fmt.Printf("etf: %s;", etf)
			c.String(http.StatusOK, "etf: %s;", etf)
		})

		// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
		v1.GET("/welcome", func(c *gin.Context) {
			// firstname := c.DefaultQuery("firstname", "Guest")
			// lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

			// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

			user := &JsonResult{Code: 0, Msg: "张三"}
			// Gin 会自动将User转换成JSon返回给前端，其他像Map,slice也一样
			c.JSON(http.StatusOK, user)

			// msg, _ := json.Marshal(JsonResult{Code: 400, Msg: "验证失败"})
			// c.JSON(http.StatusBadRequest, msg)
		})
	}

	router.Run(":7001")
}

// arkk_etf where name=TSLA
func GetStockData(c *gin.Context) {
	cond := c.Param("stock")
	db := c.Param("db")

	query := fmt.Sprintf("%s where ark_stock_name='%s'", db, cond)
	fmt.Print(query)
	result := get_data(query)

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}

func GetETFData(c *gin.Context) {
	db := c.Param("db")
	time := c.Param("time")

	if len(time) == 8 {
		time = time[4:6] + "/" + time[6:8] + "/" + time[0:4]

		query := fmt.Sprintf("%s where ark_date='%s'", db, time)
		// fmt.Print(query)
		result := get_data(query)

		c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
	} else {
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "data": "参数错误"})
	}
}

type ARK_ETF_STOCKCHANGE struct {
	Ark_Date  string
	Ark_Share string
}

func CalcStockChange(begin string, end string) string {
	nbegin, _ := strconv.ParseInt(begin, 10, 64)
	nend, _ := strconv.ParseInt(end, 10, 64)
	ret := strconv.FormatInt(nend-nbegin, 10)
	return ret
}

func GetETFStockChange(c *gin.Context) {
	db := c.Param("db")
	stock := c.Param("stock")

	query := fmt.Sprintf("%s where ark_stock_name='%s'", db, stock)
	result := get_data_count(query)

	ret := make([]ARK_ETF_STOCKCHANGE, 0)
	// calc and modify data
	//var beginDate string
	var beginShare string
	for index, data := range result {
		if index == 0 {
			//beginDate = data.Ark_Date

			beginShare = data.Ark_Shares
			continue
		}
		var ark_stock ARK_ETF_STOCKCHANGE

		ark_stock.Ark_Date = data.Ark_Date
		ark_stock.Ark_Share = CalcStockChange(beginShare, data.Ark_Shares)

		beginShare = data.Ark_Shares

		ret = append(ret, ark_stock)
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
}

func GetETFNewImport(c *gin.Context) {
	result := &JsonResult{Code: -1, Msg: "接口实现中..."}
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}

type ARK_ETF_ALL_STOCKCHANGE struct {
	Ark_Date  string
	Ark_Stock string
	Ark_Share string
}

func GetETFAllStockChange(c *gin.Context) {
	// 获取db所有数据
	db := c.Param("db")

	query := fmt.Sprintf("%s", db)
	result := get_data_count(query)

	// todo 筛选出 code: {date,change}
	// for index, data := range result {
	// 	if index == 0 {
	// 		//beginDate = data.Ark_Date
	// 		beginShare = data.Ark_Shares
	// 		continue
	// 	}
	// 	var ark_stock ARK_ETF_ALL_STOCKCHANGE

	// 	ark_stock.Ark_Date = data.Ark_Date
	// 	ark_stock.Ark_Share = CalcStockChange(beginShare, data.Ark_Shares)

	// 	beginShare = data.Ark_Shares

	// 	ret = append(ret, ark_stock)
	// }

	//result := &JsonResult{Code: -1, Msg: "接口实现中..."}
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}

func GetETFNewExport(c *gin.Context) {
	redisCLi, err := ProduceRedis(redis_host, redis_port, redis_passwd, 0, 100, true)
	if err != nil {
		fmt.Println("redis链接错误！err>>>", err.Error())
		return
	}

	// 获取全部keys
	allKeysLst := redisCLi.GetAllKeys()
	fmt.Print("key>>> ", allKeysLst)
	result := &JsonResult{Code: -1, Msg: "接口实现中..."}
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}
