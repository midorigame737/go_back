package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	engine.GET("/", func(c *gin.Context) { //多分.com直下にアクセスされたときの処理
		c.JSON(http.StatusOK, gin.H{ //多分第一引数が結果、第２引数がレスポンス内容
			"message":    "hello world",
			"User-Agent": ua,
		})
	})
	engine.Run(":3000") //ポート指定
}
