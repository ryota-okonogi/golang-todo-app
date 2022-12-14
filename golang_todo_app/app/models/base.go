package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"golang-todo-app/golang_todo_app/config"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB //変数Dbをポインタ型でグローバルに宣言

var err error //エラーも変数として定義しておく

/*
const (
	tableNameUser    = "users" //constで定数としてテーブル名を宣言する
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

// init関数 => テーブルはmain関数の前に作成する
func init() {

	url := os.Getenv("DATABASE_URL")                        //herokuの環境変数の値を取り出す事ができる //hrokuのURL("DATABASE_URL")が、変数urlに入る
	conneciton, _ := pq.ParseURL(url)                       //返り値connectionに接続する変数urlのリソースを、connectionとして取得する事ができる。
	conneciton += "sslmode=require"                         //connecitonという文字列に "sslmode=require" を足す
	Db, err = sql.Open(config.Config.SQLDriver, conneciton) //config.iniのsqlドライバを取得して渡す //リソースは取得したconnecitonを指定する
	if err != nil {
		log.Fatalln(err)
	}

	/*
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

		//Todo作成のコマンド
		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				content TEXT,
				user_id INTEGER,
				created_at DATETIME)`, tableNameTodo)
		Db.Exec(cmdT)

		//Session作成のコマンド
		cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			email STRING,
			user_id INTEGER,
			created_at DATETIME)`, tableNameSession)
		Db.Exec(cmdS)
	*/
}

//uuidとpasswordを作成する関数を用意する

// uuidを作成する関数
func createUUID() (uuidobj uuid.UUID) { //func 関数名(引数なし) (返り値 返り値の型[uuidパッケージのUUID型を使っているという意])
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// passwordを作成する関数
func Encrypt(plaintext string) (cryptext string) { //func 関数名(引数 引数の型) (返り値 返り値の型)
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
