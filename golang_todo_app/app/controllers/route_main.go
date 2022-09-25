package controllers

import (
	"net/http"
)

// layout.htmlの表示をするハンドラー
func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r) //sessionでcookieを取得する => そのcookieが「DBに存在するか」を判定する
	if err != nil {         //エラー => ログインしていない ==> generateHTMLの処理に飛ばす
		generateHTML(w, "Hello", "layout", "public_navbar", "top") //メソッド(args1: ResponseWriter, args2: data[明示的に渡している], args3: template[template1, template2])
	} else {
		http.Redirect(w, r, "/todos", 302) //ログインしている場合 => 「(/todos)のURLにアクセスする」という風にする
	}
}

// index.htmlの表示をするハンドラー
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r) //sessionを使って「ログインしているかどうか」の判定を取得する
	if err != nil {         //エラーが有る場合は
		http.Redirect(w, r, "/", 302) //ログインしていない = Topページにリダイレクトする
	} else { //ただし、セッションが存在する場合は
		generateHTML(w, nil, "layout", "private_navbar", "index") //indexを表示する
	}
}
