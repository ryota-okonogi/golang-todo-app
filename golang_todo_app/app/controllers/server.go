package controllers

import (
	"golang-todo-app/golang_todo_app/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)                               //関数名(第１引数：URL, 第２引数：ハンドラー)
	return http.ListenAndServe(":"+config.Config.Port, nil) //(パッケージ名.変数名.フィールド)
}
