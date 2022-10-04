package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	/*ginEnginの初期化
	GET,POSTとかのエンドポイントの登録してくれる
	*/
	db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/TestMirai")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	log.Print("Success Opendb")
	defer db.Close()
	ua := ""
	engine.Use(func(c *gin.Context) { //Useでミドルウェア登録,Useの引数に関数名入れて別個で関数宣言でも行ける
		ua = c.GetHeader("User-Agent")
		c.Next() //ここでUseで登録した関数実行
	})
	engine.GET("/", func(c *gin.Context) { //エンドポイント(URI)登録
		c.JSON(http.StatusOK, gin.H{ //多分第一引数が結果、第２引数がレスポンス内容
			"message":    "hello world",
			"User-Agent": ua,
		})
	})
	engine.Run(":3000") //ポート指定して鯖建てる
}
