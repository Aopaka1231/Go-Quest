package routers

import (
	"go-quest/cruds"
	"go-quest/db"
	"go-quest/schema"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TodoRouterを初期化する関数
func InitTodoRouters(tr *gin.RouterGroup) {
	//""というパスに対してGETメソッドで、getTodos関数を設定する
	tr.GET("", getTodos)
	//これも同様
	tr.POST("", postTodo)
	//"/:id"とすることで任意の文字列をidが拾ってこれるようになる
	tr.DELETE("/:id", deleteTodo)
	tr.PATCH("/:id", UpdateTodo)
}

// データベースからTodoの配列を受け取りそれをフロントに返す関数
func getTodos(c *gin.Context) {
	var (
		todos []db.Todo
		err   error
	)

	//crudsで定義したGetAllTodosを使用
	if todos, err = cruds.GetAllTodos(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//errorがない場合todosを返す
	c.JSON(http.StatusOK, &todos)
}

// フロントから送られてきた情報をデータベースに保存する関数
func postTodo(c *gin.Context) {
	var (
		err     error
		payload schema.CreateTodo
		todo    db.Todo
	)

	//BindJSONは構造体とjsonを結び付けてくれる
	//schema.CreateTodoのContentに`json:"content"`
	//と書いていたのは、フロント送られてくる{"content":"テキストテキスト"}
	//をpayloadのContent要素に勝手に代入してくれる。
	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if todo, err = cruds.CreateTodo(payload.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//errorがなければ登録したtodoを返す
	c.JSON(http.StatusOK, &todo)
}

func deleteTodo(c *gin.Context) {
	//c.Param("id")として、パスパラメータを取得できる
	//でもこれは文字列なのでuint型に変換したい
	//strconvパッケージを使用して10進数で、64bitのuintに変換
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//crudsで定義したDeleteTodoを実行
	//引数にはuintでキャストしたidを渡す
	err = cruds.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	//errorがない場合{"message": "ok!!"}を返す。
	c.JSON(http.StatusOK, gin.H{
		"message": "OK!!",
	})
}

func UpdateTodo(c *gin.Context) {
	var (
		err     error
		payload schema.CreateTodo
		todo    db.Todo
	)

	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := c.Param("id")
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if todo, err = cruds.UpdateTodo(uint(parseId), payload.Content); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	//errorがなければ登録したtodoを返す
	c.JSON(http.StatusOK, &todo)
}
