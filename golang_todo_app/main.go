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
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	//関数GetUserの実行(パッケージ名.関数名)
	u, _ := models.GetUser(1) //idが(1)番のユーザーを取得する

	fmt.Println(u)
}
