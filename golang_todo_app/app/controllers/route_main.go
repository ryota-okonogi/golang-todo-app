package controllers

import (
	"log"
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
	sess, err := session(w, r) //sessionを使って「ログインしているかどうか」の判定を取得する
	if err != nil {            //エラーが有る場合は
		http.Redirect(w, r, "/", 302) //ログインしていない = Topページにリダイレクトする
	} else { //ただし、セッションが存在する場合は
		user, err := sess.GetUserBySession() //sessionのuser_idを使って、そのuser_idと一致するuserを取得する。
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()                          //userが作成したTodoの一覧を取得する
		user.Todos = todos                                         //users.goのstructに定義した構造体をtodosに渡す
		generateHTML(w, user, "layout", "private_navbar", "index") //indexを表示する
	}
}

// todoを作成するハンドラー
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r) //ログインを確認する
	if err != nil {         //エラーだった場合
		http.Redirect(w, r, "/login", 302) //ログイン画面に遷移させる
	} else { //もしログインしている場合は
		generateHTML(w, nil, "layout", "private_navbar", "todo_new") //テンプレートを渡す
	}
}

// todo_new.htmlに記載のURL("/todos/save")に対応するハンドラーを作成する
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r) //セッションの確認
	if err != nil {            //ログインしていない場合は
		http.Redirect(w, r, "/login", 302) //ログインページにリダイレクトする
	} else { //ログインしている場合は
		err = r.ParseForm() //エラーハンドリングも含め、フォームの値を取得する。
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")            //フォームの値を取得
		if err := user.CreateTodo(content); err != nil { //取得したcontentをCreateTodoに渡してエラーハンドリングもする
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302) //最後に(/todos)にリダイレクトさせる
	}
}
