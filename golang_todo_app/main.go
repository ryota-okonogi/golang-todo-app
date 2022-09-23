package main

import (
	"fmt"
	"golang-todo-app/golang_todo_app/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port) //(パッケージ名.変数名.フィールド)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	//init関数を呼び出す処理(意味は無いらしい)
	fmt.Println(models.Db) //パッケージ名.変数名

	//CreateUserのメソッドを作成したら、こちらでテストをする。

	/*
		//user定義
		u := &models.User{}
		u.Name = "test2"
		u.Email = "test2@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*

		//関数GetUserの実行(パッケージ名.関数名)
		u, _ := models.GetUser(1) //idが(1)番のユーザーを取得する

		fmt.Println(u)

		//GetUserで取得したユーザーを更新する
		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1) //GetUserを使ってuserを表示する
		fmt.Println(u)

		//GetUserで取得したユーザーを削除する
		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

	*/

	/*
		//todosテーブルの作成
		user, _ := models.GetUser(2) //idが(2)番のユーザーを取得する
		user.CreateTodo("First Todo")
	*/

	/*
		//関数GetTodoの実行(パッケージ名.関数名)
		t, _ := models.GetTodo(1) //CreateTodoで作成した "First Todo" を取得する
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(3) //idが(3)番のユーザーを取得する
		user.CreateTodo("Third Todo")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
}
