package modules

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-api/db"
	"net/http"
)

type data struct {
	ID      string    `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func GetDatas(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, title, message FROM reminder.reminder")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var datas []data
	for rows.Next() {
		var a data
		err := rows.Scan(&a.ID, &a.Title, &a.Message)
		if err != nil {
			log.Fatal(err)
		}
		datas = append(datas, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, datas)
}
