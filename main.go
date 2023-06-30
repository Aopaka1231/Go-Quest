// packageというものをmainという名前で定義
package main

import (
	"go-quest/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//ginを初期化してrという変数に代入
	//goでは、:=とすることで暗黙的に型宣言を行っている
	r := gin.Default()

	db.InitDB()

	//GETメソッドを定義
	//"/"にアクセスしたときに{"message": "Hello world!!}を返す
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!!",
		})
	})

	//ginを実行する
	//()内に下のような書き方をするとポート番号を指定することができる
	r.Run(":8000")
}
