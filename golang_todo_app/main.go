package main

import (
	"fmt"
	"golang-todo-app/golang_todo_app/app/controllers"
	"golang-todo-app/golang_todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
