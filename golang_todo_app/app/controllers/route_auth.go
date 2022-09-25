package controllers

import (
	"golang-todo-app/golang_todo_app/app/models"
	"log"
	"net/http"
)

// "sample_todo/app/models"

// signupのハンドラー
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //r.Method とする事で、右辺で指定したリクエストのメソッドを取得する事ができる
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
		//入力フォームの解析(r.Method == "POST")
	} else if r.Method == "POST" { //Postの場合で行う処理 = 入力フォームで入力された値を元に新しいユーザーを作成する
		err := r.ParseForm() //r.ParseFormとする事で、「入力フォームの解析」を行う。
		if err != nil {      //エラーハンドリング
			log.Println(err)
		}
		//まずは「入力フォームの値」を取得する
		//入力された値を受け取ってユーザーを作成したい為、「userのストラクタを作成」する。
		user := models.User{
			Name:     r.PostFormValue("name"), //r.PostFormValue = signup.htmlのimputタグで指定されている属性から「値を取得」する事ができる。引数として属性のnameを渡す。
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil { //ユーザー作成のメソッド
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302) //ユーザー登録が成功したらTopページにリダイレクトさせたい => http.Redirect(w, r, "/", 302) = http.Redirect(ResponseWriter, *http.Request, "リダイレクト先のURL", ステータスコード)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// loginのハンドラー
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()                                         //フォームから値を取得する
	user, err := models.GetUserByEmail(r.PostFormValue("email")) //フォームのEmailからユーザーを取得する //(r.PostFormValue("email")) => login.htmlの name="email" を指している
	if err != nil {                                              //エラーハンドリング
		log.Println(err)                   //もしエラーの場合は
		http.Redirect(w, r, "/login", 302) //ログインできない => httpのリダイレクトでloginページにリダイレクトさせる
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) { //「ユーザーが存在してパスワードがフォームで入力されたpasswordと一致する場合はログインできる」ことを判定する
		session, err := user.CreateSession() //passwordが一致したuserでsessionを作成する
		if err != nil {                      //エラーハンドリング
			log.Println(err)
		}

		//以下はパターンとして覚える
		cookie := http.Cookie{ //http.Cookie = struct
			Name:     "_cookie",
			Value:    session.UUID, //CreateSessionで作成されたsessionのUUID
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302) //現時点ではログインページを作成していない為、Topページにリダイレクトさせる。
	} else { //パスワードが一致しない場合
		http.Redirect(w, r, "/login", 302)
	}
}

/*
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)
	}
*/

/*
		switch r.Method {
		case http.MethodGet:
			_, err := session(w, r)
			if err != nil {
				generateHTML(w, nil, "layout", "public_navbar", "signup")
			} else {
				http.Redirect(w, r, "/todos", 302)
			}
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
			}
			user := models.User{
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				PassWord: r.PostFormValue("password"),
			}
			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/", 302)
		}
	}

	func login(w http.ResponseWriter, r *http.Request) {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "login")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	}

	func authenticate(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)

		}
		user, err := models.GetUserByEmail(r.PostFormValue("email"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", 302)
		}
		if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				log.Println(err)
			}
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}

	func logout(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("_cookie")
		if err != nil {
			log.Println(err)
		}
		if err != http.ErrNoCookie {
			session := models.Session{UUID: cookie.Value}
			session.DeleteSessionByUUID()
		}
		http.Redirect(writer, request, "/login", 302)
*/
