package models

import (
	"database/sql"
	"fmt"
	"golang-todo-app/golang_todo_app/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB //変数Dbをポインタ型でグローバルに宣言

var err error //エラーも変数として定義しておく

const (
	tableNameUser = "users" //constで定数としてテーブル名を宣言する
)

// init関数 => テーブルはmain関数の前に作成する
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName) //(パッケージ名.変数名.フィールド)
	if err != nil {
		log.Fatalln(err)
	}

	//user作成のコマンド
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)
}
