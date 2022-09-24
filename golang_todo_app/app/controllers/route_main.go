package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// ハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("golang_todo_app/app/views/templates/top.html") //パッケージ名.関数名
	if err != nil {
		log.Fatalln(err) //エラーハンドリング
	}
	t.Execute(w, "Hello") //tの実行(args1: http.ResponseWriter, args2: データ[今回はnil])
}
