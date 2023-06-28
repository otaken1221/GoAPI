package main

import (

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"go-api/modules"
)

func main() {
	router := gin.Default()
	router.GET("/datas", modules.GetDatas)
	router.Run("localhost:8080")
}

