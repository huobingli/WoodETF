package page

import (
	"ZEUS/server/hxapp/dao"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// 获取所有页面数据
func GetAllPage(c *gin.Context) {
	db, err := sql.Open("sqlite3", "C:\\Users\\huobingli\\go\\src\\Zeus\\export_sav.db")
	dao.CheckErr(err)
	rows, err := db.Query("SELECT ID,NAME FROM CDrawPage")
	ret := make([]dao.HX_Draw, 0)
	if rows != nil {
		for rows.Next() {
			var hxdraw dao.HX_Draw
			if err := rows.Scan(&hxdraw.HX_ID, &hxdraw.HX_NAME); err != nil {
				log.Fatal(err)
			}

			ret = append(ret, hxdraw)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
}

// 查某个页面所有容器
func GetAllPageContainer(c *gin.Context) {
	pageid := c.Param("pageid")
	db, err := sql.Open("sqlite3", dao.Sql_file)
	dao.CheckErr(err)
	query := fmt.Sprintf("select ID,NAME from IDs where ID in ( SELECT m_dwID FROM CDrawPage_CDrawContainer where PID=%s)", pageid)
	rows, err := db.Query(query)
	ret := make([]dao.HX_Draw, 0)
	if rows != nil {
		for rows.Next() {
			var hxdraw dao.HX_Draw
			if err := rows.Scan(&hxdraw.HX_ID, &hxdraw.HX_NAME); err != nil {
				log.Fatal(err)
			}

			ret = append(ret, hxdraw)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
}


func GetPageHaveContainer(c *gin.Context) {
	contid := c.Param("contid")
	db, err := sql.Open("sqlite3", "C:\\Users\\huobingli\\go\\src\\Zeus\\export_sav.db")
	dao.CheckErr(err)
	//query := fmt.Sprintf("SELECT PID  FROM CDrawPage_CDrawContainer where m_dwID=%s", contid)
	query := fmt.Sprintf("select ID,NAME from IDs where ID in ( SELECT PID  FROM CDrawPage_CDrawContainer where m_dwID=%s)", contid)
	rows, err := db.Query(query)
	ret := make([]dao.HX_Draw, 0)
	for rows.Next() {
		var hxdraw dao.HX_Draw
		if err := rows.Scan(&hxdraw.HX_ID, &hxdraw.HX_NAME); err != nil {
			log.Fatal(err)
		}

		ret = append(ret, hxdraw)
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
}