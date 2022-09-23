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

// Todoを取得する関数
func GetTodo(id int) (todo Todo, err error) { //func 関数名(引数 引数の型) (返り値1 型, 返り値2 型)
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan( //scan データ追加
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}

// 複数のTodoを取得する関数(Todoのリストを取得する)
func GetTodos() (todos []Todo, err error) { //func 関数名(引数なし) (返り値1 型, 返り値2 型)
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd) //Query は条件に合うものを全て取得
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() { // for rows.Next() => 取得したデータをループでスライスに追加
		var todo Todo //変数宣言
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil { //エラーハンドリング
			log.Fatalln(err)
		}
		todos = append(todos, todo) //todosのリストにappendしていく => 変数名 = append(追加したいスライス, 追加するデータ)
	}
	rows.Close() //rowsの処理を終わらせる

	return todos, err
}

// 「特定のユーザーの Todoのリストを取得する」という関数
func (u *User) GetTodosByUser() (todos []Todo, err error) { //func (レシーバーの名前 型) 関数名(引数なし) (返り値1 型, 返り値2 型) {処理内容}
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo //変数宣言
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil { //エラーハンドリング
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close() //rowsの処理を終わらせる

	return todos, err
}
