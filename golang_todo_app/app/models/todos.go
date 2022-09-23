package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) { //func (レシーバーの名前 型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now()) //Dbは、users.goファイルには存在しないが、modelsパッケージに存在する為「パッケージ名」を指定しなくとも使用する事ができる。
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// passwordを作成する関数
func GetTodo(id int) (todo Todo, err error) { //func 関数名(引数 引数の型) (返り値1 型, 返り値2 型)
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}
