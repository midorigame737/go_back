package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	/*ginEnginの初期化
	GET,POSTとかのエンドポイントの登録してくれる
	*/
	ua := ""
	engine.Use(func(c *gin.Context) { //Useでミドルウェア登録
		ua = c.GetHeader("User-Agent")
		c.Next() //ここでUseで登録した関数実行
	})
	engine.GET("/", func(c *gin.Context) { //多分.com直下にアクセスされたときの処理
		c.JSON(http.StatusOK, gin.H{ //多分第一引数が結果、第２引数がレスポンス内容
			"message":    "hello world",
			"User-Agent": ua,
		})
	})
	engine.Run(":3000") //ポート指定
}
