package controllers

import (
	"golang-todo-app/golang_todo_app/config"
	"net/http"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files)) //args1: URLPath, args2: files{ ("/static", files)にするのはstaticをPathの階層から取り除くため }

	http.HandleFunc("/", top)                               //関数名(第１引数：URL, 第２引数：ハンドラー)
	return http.ListenAndServe(":"+config.Config.Port, nil) //(パッケージ名.変数名.フィールド)
}
