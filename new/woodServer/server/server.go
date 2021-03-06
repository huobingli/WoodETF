// +build windows

package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"

	"woodServer/server/router/woodapp"
)

var (
	g errgroup.Group
)

func router_woodetf() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router_test() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

//启动
func main() {
	//router := gin.New()
	router := gin.Default()

	// 设置静态资源
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 分组
	//v1 := router.Group("/v1")
	//{
	//	v1.GET("/GetStockData/:db/:stock", GetStockData)
	//	v1.GET("/GetETFData/:db/:time", GetETFData)
	//
	//	// 获取某etf成分股变化 8.30起
	//	v1.GET("/GetETFStockChange/:db/:stock", GetETFStockChange)
	//
	//	// 获取某etf新进成分股数据
	//	v1.GET("GetETFNewImportStock/:db/:time", GetETFNewImport)
	//	v1.GET("GetETFNewExportStock/:db/:time", GetETFNewExport)
	//
	//	// 获取某etf 所有成分股变化
	//	v1.GET("GetETFAllStockChange/:db/", GetETFAllStockChange)
	//
	//	v1.POST("/GetStockData", func(c *gin.Context) {
	//
	//		stock := c.Query("stock")
	//
	//		// 获取StockData
	//		fmt.Printf("stock: %s;", stock)
	//		c.String(http.StatusOK, "stock: %s;", stock)
	//	})
	//
	//	// v1.POST("/GetETFData/:db/:time", GetETFData)
	//	v1.POST("/GetETFData", func(c *gin.Context) {
	//
	//		etf := c.Query("etf")
	//
	//		// fmt.Printf("etf: %s;", etf)
	//		c.String(http.StatusOK, "etf: %s;", etf)
	//	})
	//
	//	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	//	v1.GET("/welcome", func(c *gin.Context) {
	//		// firstname := c.DefaultQuery("firstname", "Guest")
	//		// lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
	//
	//		// c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//
	//		user := &test2.JsonResult{Code: 0, Msg: "张三"}
	//		// Gin 会自动将User转换成JSon返回给前端，其他像Map,slice也一样
	//		c.JSON(http.StatusOK, user)
	//
	//		// msg, _ := json.Marshal(JsonResult{Code: 400, Msg: "验证失败"})
	//		// c.JSON(http.StatusBadRequest, msg)
	//	})
	//}

	//home := router.Group("/home")
	//{
	//	home.GET("/index/:db/:time", Index)
	//}

	//
	// gin 路由重定向
	router.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//
	//router.Use(gin.Recovery())



	server01 := &http.Server{
		Addr:         ":7003",
		Handler:      router_woodetf(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":7004",
		Handler:      router_test(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server03 := &http.Server{
		Addr:         ":7001",
		Handler:      woodapp.Wood_app_router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	g.Go(func() error {
		return server03.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	router.Run(":7001")
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": "pong"})
}

// arkk_etf where name=TSLA
//func GetStockData(c *gin.Context) {
//	cond := c.Param("stock")
//	db := c.Param("db")
//
//	query := fmt.Sprintf("%s where ark_stock_name='%s'", db, cond)
//	fmt.Print(query)
//	result := mymysql.get_data(query)
//
//	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//}
//
//func GetETFData(c *gin.Context) {
//	db := c.Param("db")
//	time := c.Param("time")
//
//	if len(time) == 8 {
//		time = time[4:6] + "/" + time[6:8] + "/" + time[0:4]
//
//		query := fmt.Sprintf("%s where ark_date='%s'", db, time)
//		// fmt.Print(query)
//		result := mymysql.get_data(query)
//
//		c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//	} else {
//		c.JSON(http.StatusOK, gin.H{"status_code": -1, "data": "参数错误"})
//	}
//}
//
//type ARK_ETF_STOCKCHANGE struct {
//	Ark_Date  string
//	Ark_Share string
//}
//
//func CalcStockChange(begin string, end string) string {
//	nbegin, _ := strconv.ParseInt(begin, 10, 64)
//	nend, _ := strconv.ParseInt(end, 10, 64)
//	ret := strconv.FormatInt(nend-nbegin, 10)
//	return ret
//}
//
//func GetETFStockChange(c *gin.Context) {
//	db := c.Param("db")
//	stock := c.Param("stock")
//
//	query := fmt.Sprintf("%s where ark_stock_name='%s'", db, stock)
//	result := mymysql.get_data_count(query)
//
//	ret := make([]ARK_ETF_STOCKCHANGE, 0)
//	// calc and modify data
//	//var beginDate string
//	var beginShare string
//	for index, data := range result {
//		if index == 0 {
//			//beginDate = data.Ark_Date
//
//			beginShare = data.Ark_Shares
//			continue
//		}
//		var ark_stock ARK_ETF_STOCKCHANGE
//
//		ark_stock.Ark_Date = data.Ark_Date
//		ark_stock.Ark_Share = CalcStockChange(beginShare, data.Ark_Shares)
//
//		beginShare = data.Ark_Shares
//
//		ret = append(ret, ark_stock)
//	}
//
//	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
//}
//
//func GetETFNewImport(c *gin.Context) {
//	result := &test2.JsonResult{Code: -1, Msg: "接口实现中..."}
//	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//}
//
//type ARK_ETF_ALL_STOCKCHANGE struct {
//	Ark_Date  string
//	Ark_Stock string
//	Ark_Share string
//}
//
//func GetETFAllStockChange(c *gin.Context) {
//	// 获取db所有数据
//	db := c.Param("db")
//
//	query := fmt.Sprintf("%s", db)
//	result := mymysql.get_data_count(query)
//
//	// todo 筛选出 code: {date,change}
//	// for index, data := range result {
//	// 	if index == 0 {
//	// 		//beginDate = data.Ark_Date
//	// 		beginShare = data.Ark_Shares
//	// 		continue
//	// 	}
//	// 	var ark_stock ARK_ETF_ALL_STOCKCHANGE
//
//	// 	ark_stock.Ark_Date = data.Ark_Date
//	// 	ark_stock.Ark_Share = CalcStockChange(beginShare, data.Ark_Shares)
//
//	// 	beginShare = data.Ark_Shares
//
//	// 	ret = append(ret, ark_stock)
//	// }
//
//	//result := &JsonResult{Code: -1, Msg: "接口实现中..."}
//	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//}
//
////func todata(interface{} data) string {
////
////}
//
//func GetETFNewExport(c *gin.Context) {
//	redisCLi, err := myredis.ProduceRedis(config.Redis_host, config.Redis_port, config.Redis_passwd, 0, 100, true)
//	if err != nil {
//		fmt.Println("redis链接错误！err>>>", err.Error())
//		return
//	}
//
//	db := c.Param("db")
//	time := strings.Replace(c.Param("time"), "-", "/", -1)
//	list_name := strings.Join([]string{db, strings.ReplaceAll(time, "/", "")}, "")
//
//	if redisCLi.HashExist(db, list_name) {
//		result, err := redisCLi.HashGet(db, list_name)
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{"status_code": -1, "data": err})
//			//	上报错误
//		} else {
//			//value, _ := json.Marshal(result)
//			//fmt.Println(result)
//			c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//		}
//
//	} else {
//		query := fmt.Sprintf("%s where ark_date='%s'", db, time)
//		result := mymysql.get_data_with_time(query)
//		c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//
//		value, _ := json.Marshal(result)
//		redisCLi.HashAdd(db, list_name, string(value))
//
//		allKeysLst := redisCLi.GetAllKeys()
//		fmt.Print("key>>> ", allKeysLst)
//	}
//}
//
//func Index(c *gin.Context) {
//	redisCLi, err := myredis.ProduceRedis(config.Redis_host, config.Redis_port, config.Redis_passwd, 0, 100, true)
//	if err != nil {
//		fmt.Println("redis链接错误！err>>>", err.Error())
//		return
//	}
//
//	db := c.Param("db")
//	time := strings.Replace(c.Param("time"), "-", "/", -1)
//	list_name := strings.Join([]string{db, strings.ReplaceAll(time, "/", "")}, "")
//
//	result, err := redisCLi.StringGet(list_name)
//	if err != nil {
//		if redisCLi.IsNill(err) {
//			query := fmt.Sprintf("%s where ark_date='%s'", db, time)
//			result := mymysql.get_data_with_time(query)
//			c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//
//			value, _ := json.Marshal(result)
//			redisCLi.StringSet(list_name, string(value))
//			//c.JSON(http.StatusOK, gin.H{"status_code": -1, "data": "nil data"})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"status_code": -1, "data": err})
//		}
//	} else {
//		c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": result})
//	}
//}
