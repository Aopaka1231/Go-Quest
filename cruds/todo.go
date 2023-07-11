package cruds

import (
	"go-quest/db"
)

// Todoをデータベースに保存する関数
func CreateTodo(content string) (res_todo db.Todo, err error) {
	res_todo = db.Todo{Content: content}
	err = db.Sspl.Create(&res_todo).Error
	return
}

// すべてのTodoの配列を返す関数
func GetAllTodos() (res_todos []db.Todo, err error) {
	err = db.Sspl.Find(&res_todos).Error
	return
}

// Todoを更新する関数
func UpdateTodo(id uint, content string, Todo db.Todo) (err error) {
	if err = db.Sspl.Where("id = ?", id).Error; err != nil {
		return err
	}
	db.Sspl.Model(&Todo).Where("id = ?", id).Update(content, "Content")
	return
}

// 引数のidのTodoを削除する関数
func DeleteTodo(id uint) (err error) {
	err = db.Sspl.Where("id = ?", id).Delete(&db.Todo{}).Error

	return
}
