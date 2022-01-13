package woodapp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wood_app_router() http.Handler {
	router := gin.New()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 分组
	v1 := router.Group("/v1")
	{
		v1.GET("/GetStockData/:db/:stock", GetStockData)
		v1.GET("/GetETFData/:db/:time", GetETFData)

		// 获取某etf成分股变化 8.30起
		v1.GET("/GetETFStockChange/:db/:stock", GetETFStockChange)

		// 获取某etf新进成分股数据
		v1.GET("GetETFNewImportStock/:db/:time", GetETFNewImport)
		v1.GET("GetETFNewExportStock/:db/:time", GetETFNewExport)

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

	return router
}