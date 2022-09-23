package models

import (
	"log"
	"time"
)

// 構造体の生成
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// ユーザーの作成
func (u *User) CreateUser() (err error) { //func (任意のレシーバーの名前 型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, //Dbは、users.goファイルには存在しないが、modelsパッケージに存在する為「パッケージ名」を指定しなくとも使用する事ができる。
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err

}

// メソッドではなく関数
func GetUser(id int) (user User, err error) { //func 関数名(引数 引数の型) (返り値1 型, 返り値2 型)
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}
