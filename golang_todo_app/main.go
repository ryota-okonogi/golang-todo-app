package main

import (
	"fmt"
	"golang-todo-app/golang_todo_app/app/models"
	"log"
)

func main() {
	fmt.Println(models.Db)

	// controllers.StartMainServer()

	user, _ := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	//上記で作成したuserを使ってセッションを作成して受け取り、表示する。
	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	//上記で作成したセッションを使って、CheckSessionを使う。
	valid, _ := session.CheckSession()
	fmt.Println(valid)
}
