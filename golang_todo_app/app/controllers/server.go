package controllers

import (
	"fmt"
	"golang-todo-app/golang_todo_app/app/models"
	"golang-todo-app/golang_todo_app/config"
	"html/template"
	"net/http"
)

// func 関数名(args1: ハンドラー関数の引数に設定したメソッド, args2: 型, args3: 可変調引数をstring型で設定)
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	//filenamesの中身のファイルを、スライス型の変数filesに格納していく処理
	var files []string
	for _, file := range filenames { //rangeで、引数として渡したfilenamesの中身を取り出す
		files = append(files, fmt.Sprintf("golang_todo_app/app/views/templates/%s.html", file)) //取り出したものをfmt.Sprintfを使ってファイルパスに入れて、filesに格納する。(Mustの引数としてParseFilesを渡す事で、エラーの場合は「パニック」状態になる。)
	}

	templates := template.Must(template.ParseFiles(files...)) ////Must = templateをあらかじめキャッシュしておいて効率的に処理できる様にしている
	templates.ExecuteTemplate(w, "layout", data)              //実行コマンド(args1: ResponseWriter, args2: 実行するテンプレート[明示的に渡している], args3: data)
}

// cookieを取得する関数
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) { //func 関数名(args1: 型, args2: 型) (返り値1: 型, 返り値2: 型) {処理内容}
	//httpリクエストからcookieを取得する
	cookie, err := r.Cookie("_cookie") //cookieのNameとして渡した値を指定してcookieを取得する事ができる
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok { //もし存在しない場合は
			err = fmt.Errorf("Invalid session") //エラーを生成する("Invalid session" => 無効なセッション)
		}
	}
	return sess, err //sessionとエラーを返す
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files)) //args1: URLPath, args2: files{ ("/static", files)にするのはstaticをPathの階層から取り除くため }

	http.HandleFunc("/", top) //関数名(第１引数：URL, 第２引数：ハンドラー)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)                        //ログインしているユーザーしか(/todos)にアクセスできない
	return http.ListenAndServe(":"+config.Config.Port, nil) //(パッケージ名.変数名.フィールド)
}
