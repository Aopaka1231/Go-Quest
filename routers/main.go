package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//corsミドルウェアを使用する
	//おまじない
	r.Use(cors.New(cors.Config{
		//アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		//アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		//許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
	}))

	//TodoRouterのすべてのパスを"/todo"で始まるようにグループ化
	todo_router := r.Group("/todo")
	//InitTodoRouterを実行
	InitTodoRouters(todo_router)
}
