package controllers

import (
	"net/http"
)

// ハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "public_navbar", "top") //メソッド(args1: ResponseWriter, args2: data[明示的に渡している], args3: template[template1, template2])
}
