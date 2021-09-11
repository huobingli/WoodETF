package main

import (
	"fmt"
	"net/http"

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
		v1.GET("/GetETFStockChange/:db/:stock", GetETFStockChange) //	等同于第一个，用于获取某个代码所有数据

		v1.GET("GetETFNewImportStock/:db/:time", GetETFNewImport)

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

func GetETFStockChange(c *gin.Context) {
	db := c.Param("db")
	stock := c.Param("stock")

	query := fmt.Sprintf("%s where ark_stock_name='%s'", db, stock)
	result := get_data_count(query)

	// todo calc change
	for _, data := range result {
		// for it, subdata := range data {
		// 	fmt.Print(it)
		// 	fmt.Print(subdata)
		// }
		fmt.Print(data.Ark_Date)
		fmt.Print(data.Ark_Shares)
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}

func GetETFNewImport(c *gin.Context) {
	result := &JsonResult{Code: -1, Msg: "接口实现中..."}
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}

func GetETFNewExport(c *gin.Context) {
	result := &JsonResult{Code: -1, Msg: "接口实现中..."}
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
}
