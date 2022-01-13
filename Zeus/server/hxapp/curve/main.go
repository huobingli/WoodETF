package curve

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"ZEUS/server/hxapp/dao"

	_ "github.com/mattn/go-sqlite3"
)

func GetAllCurve(c *gin.Context) {
	db, err := sql.Open("sqlite3", "D:\\go\\gopath\\Zeus\\export_sav.db")
	dao.CheckErr(err)
	rows, err := db.Query("SELECT ID,NAME FROM CDrawCurve")
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
