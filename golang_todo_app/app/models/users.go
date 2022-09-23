package models

import (
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

/*

// ユーザーの作成

// メソッド
func (u *User) CreateUser() (err error) { //func (レシーバーの名前 型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
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

//ユーザー情報の取得

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

//ユーザー情報の更新

// メソッド
func (u *User) UpdateUser() (err error) { //func (レシーバーの名前 型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
	cmd := `update users set name = ?, email = ? where id = ?` //update users で指定したidの name, emailを更新するというコマンド
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザーの削除

// メソッド
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?` //delete from users で指定したidの user を削除するというコマンド
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}
*/
