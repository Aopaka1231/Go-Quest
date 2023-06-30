package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Sspl *gorm.DB
)

// データベースを初期化する関数の定義
func InitDB() (err error) {
	//ここでsqliteを開く
	Sspl, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//エラーハンドリング
	if err != nil {
		return
	}

	//さっき作ったTodoモデルでテーブルを作成する
	if err = Sspl.AutoMigrate(&Todo{}); err != nil {
		return
	}

	return
}
