package hxapp

import (
	"ZEUS/server/hxapp/container"
	"ZEUS/server/hxapp/curve"
	"ZEUS/server/hxapp/more"
	"ZEUS/server/hxapp/page"
	"ZEUS/server/hxapp/table"
	"ZEUS/server/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hx_app_router() http.Handler {
	router := gin.New()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Use(middleware.CorsMiddle())
	v1help := router.Group("/v1/Help")
	{
		v1help.GET("/", HelpInfo)
	}

	// 分组
	v1page := router.Group("/v1/Page")
	{
		v1page.GET("/GetAllPage", page.GetAllPage)
		v1page.GET("/GetAllPageContainer/:pageid", page.GetAllPageContainer)
		v1page.GET("/GetPageHaveContainer/:contid", page.GetPageHaveContainer)
	}

	v1container := router.Group("/v1/Container")
	{
		v1container.GET("/GetAllContainer", container.GetAllContainer)
	}

	v1curve := router.Group("/v1/Curve")
	{
		v1curve.GET("/GetAllCurve", curve.GetAllCurve)
	}

	v1table := router.Group("/v1/Table")
	{
		v1table.GET("/GetAllTable", table.GetAllTable)
	}

	v1more := router.Group("/v1/More")
	{
		v1more.GET("/GetDrawIDInfo/:drawid", more.GetDrawIDInfo)
	}

	return router
}

func HelpInfo(c *gin.Context) {
	ret := `
	page：
	/v1/Page/GetAllPage	获取所有页面
	/v1/Page/GetAllPageContainer/:pageid	页面中所有容器 
	/v1/Page/GetPageHaveContainer/:contid	容器在那些页面中

	curve：
	/v1/Curve/GetAllCurve	获取所有曲线

	table：
	/v1/Table/GetAllTable	获取所有表格

	more：
	/v1/More/GetDrawIDInfo	获取ID详细信息
	`

	fmt.Fprintln(c.Writer, ret)
}