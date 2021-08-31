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
		v1.GET("/getstockdata/:stock", GetStockData)
		v1.POST("/GetStockData", func(c *gin.Context) {

			id := c.Query("id")
			page := c.DefaultQuery("page", "0")
			name := c.PostForm("name")
			message := c.PostForm("message")

			fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		})

		// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
	}

	router.Run(":7001")
}

func GetStockData(c *gin.Context) {
	url := c.Param("stock")

	redis_get

	c.JSON(http.StatusOK, url)
}
