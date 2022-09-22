package main

import (
	"fmt"
	"golang-todo-app/golang_todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port) //(パッケージ名.変数名.フィールド)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
}
