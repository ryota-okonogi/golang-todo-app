package controllers

import (
	"golang-todo-app/golang_todo_app/app/models"
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

// todoの編集をするハンドラー(Edit)
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil { //エラーの場合
		http.Redirect(w, r, "/login", 302) //ログイン画面にリダイレクトする
	} else {
		_, err := sess.GetUserBySession() //ここではuserは使わない為、返り値1は無しに設定する。
		if err != nil {                   //userが居ない場合はエラーが返ってくる(エラーハンドリング)
			log.Println(err)
		}
		t, err := models.GetTodo(id) //引数のIDからTodoを取得する
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit") //generateHTML(ResponseWriter, 返り値t, テンプレート1, テンプレート2)
	}
}

// todoの更新をするハンドラー(Update)
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil { //エラーの場合
		http.Redirect(w, r, "/login", 302)
	} else { //セッションがある場合
		err := r.ParseForm() //エラーハンドリングを含め、フォームの値を取得する
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession() //userを取得する
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content") //フォームから値を取得する(index.htmlに記載の "content" を指定する)
		//Todoのstructを作る
		t := &models.Todo{ID: id, Content: content, UserID: user.ID} //ID = 引数のid, Content = 取得したcontent, UserID = セッションから取得したusr.ID
		if err := t.UpdateTodo(); err != nil {                       //エラーハンドリングも同時にやる //structの「 t 」の UpdateTodoメソッドを使う
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302) //Updateしたら(/todos)にリダイレクトする
	}
}

// todoの削除をするハンドラー(Delete)
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r) //セッションを確認する
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession() //セッションがあればuserを確認する(sess.GetUserBySession)
		if err != nil {                   //エラーハンドリング
			log.Println(err)
		}
		t, err := models.GetTodo(id) //idからtodoを取得する(models.GetTodo)
		if err != nil {              //エラーハンドリング
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil { //(セッションがある場合[else]の)エラーハンドリング => t.DeleteTodo
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302) //削除したらhttp.Redirectで(/todos)にリダイレクトさせる
	}
}
