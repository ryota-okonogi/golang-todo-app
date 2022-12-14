package controllers

import (
	"fmt"
	"golang-todo-app/golang_todo_app/app/models"
	"golang-todo-app/golang_todo_app/config"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$") //URLの正規表現のパターンをコンパイルする

// リクエストがあったら、そのURLからIDを取得する関数
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc { //func 関数名(args略称 func(args1, args2, args3[最後のPathをint型として受け取る])) 返り値の型 {処理内容}
	return func(w http.ResponseWriter, r *http.Request) { //http.HandlerFunc(w, r)を返す
		// /todos/edit/1           (この様なURLのパターンから最後のIDを受け取って処理する)
		q := validPath.FindStringSubmatch(r.URL.Path) //URLパス(r.URL.Path) とvalidPathで登録されたURL(regexp.MustCompile〜)がマッチした部分をqでスライスとして取得する
		if q == nil {                                 //何もマッチしない場合は(スライスは空になる為)
			http.NotFound(w, r) //NotFoundを返す
			return
		}
		qi, err := strconv.Atoi(q[2]) //(マッチしたら) スライスqのインデックス番号[2]は変数validPathの(/todos/(edit|update)/([0-9])部分のIDが入っているとして、もし数値型ならAtoiで変換できる為、返り値1のqiで数値型として値を受け取る。
		if err != nil {               //数値型でない場合 = エラー(となるので)
			http.NotFound(w, r) //NotFoundを返す
			return
		}

		fn(w, r, qi) //数値型の場合 = 関数fnに引数で渡した関数(w, r, qi)を実行してIDを渡す
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static", files)) //args1: URLPath, args2: files{ ("/static", files)にするのはstaticをPathの階層から取り除くため }

	http.HandleFunc("/", top) //関数名(第１引数：URL, 第２引数：ハンドラー)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index) //ログインしているユーザーしか(/todos)にアクセスできない
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit)) //http.HandleFunc("Path", 関数(関数の引数))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	port := os.Getenv("PORT")                 //PORTを取得する
	return http.ListenAndServe(":"+port, nil) //(パッケージ名.変数名.フィールド) //取得したPORTを引数として渡す
}
