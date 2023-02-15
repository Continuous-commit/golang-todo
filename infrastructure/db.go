package infrastructure

import (
	"fmt"
	"go-todo/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DbInit() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(mysql_db)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("could not open database")
	}
	db.AutoMigrate(&domain.Todo{})
	return db
}

func DbCreate(todo domain.Todo) {
	db, err := gorm.Open("mysql", "user:password@tcp(mysql_db)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("could not open database")
	}
	db.Create(&todo)
}

func DbRead(id ...int) []domain.Todo {
	db, err := gorm.Open("mysql", "user:password@tcp(mysql_db)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("could not open database")
	}
	var todos []domain.Todo
	db.Find(&todos)
	return todos
}

func DbUpdate(id int, text string, status domain.Status, deadline string) domain.Todo {
	db, err := gorm.Open("mysql", "user:password@tcp(mysql_db)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	todo.Deadline = deadline
	db.Save(&todo)
	return todo
}

func DbDelete(id int) {
	db, err := gorm.Open("mysql", "user:password@tcp(mysql_db)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("could not open database")
	}
	var todo domain.Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
