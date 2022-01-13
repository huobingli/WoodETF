package more

import (
	"ZEUS/server/hxapp/dao"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetDrawIDInfo(c *gin.Context) {
	drawid := c.Param("drawid")
	db, err := sql.Open("sqlite3", "D:\\go\\gopath\\Zeus\\export_sav.db")
	dao.CheckErr(err)
	//query := fmt.Sprintf("SELECT PID  FROM CDrawPage_CDrawContainer where m_dwID=%s", contid)
	query := fmt.Sprintf("select ID,NAME from IDs where ID = %s", drawid)
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