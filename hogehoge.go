package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) { //多分.com直下にアクセスされたときの処理
		c.JSON(http.StatusOK, gin.H{ //多分第一引数が結果、第２引数がレスポンス内容
			"message": "hello world",
		})
	})
	engine.Run(":3000") //ポート指定
}
