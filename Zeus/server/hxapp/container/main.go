package container

import (
	"ZEUS/server/hxapp/dao"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetAllContainer(c *gin.Context) {
	db, err := sql.Open("sqlite3", dao.Sql_file)
	dao.CheckErr(err)
	rows, err := db.Query("SELECT PID,m_strProperty FROM CDrawPage_CDrawContainer")
	ret := make([]dao.HX_Container, 0)
	if rows != nil {
		for rows.Next() {
			var hxdraw dao.HX_Container
			if err := rows.Scan(&hxdraw.HX_ID, &hxdraw.HX_PROPERTY); err != nil {
				log.Fatal(err)
			}

			ret = append(ret, hxdraw)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status_code": 0, "data": ret})
}