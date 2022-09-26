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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt time.Time
}

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

// (ユーザー => Email, DB => ユーザー)情報を取得する関数
func GetUserByEmail(email string) (user User, err error) { //func 関数名(引数 引数の型) (返り値1 型, 返り値2 型)
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

// Sessionの作成をするメソッド
func (u *User) CreateSession() (session Session, err error) { //func (レシーバーの名前 型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
	session = Session{}
	cmd1 := `insert into sessions (
			uuid,
			email,
			user_id,
			created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now()) //エラーハンドリングで「作成したSessionを取得してそのまま返す」という風にする

	if err != nil {
		log.Fatalln(err)
	}

	//cmd1で作成したSessionをそのまま取得する
	cmd2 := `select id, uuid, email, user_id, created_at
	 from sessions where user_id = ? and email = ?` //usre_id と emailアドレスが一致するものを取得する

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan( //取得したデータをScanしてsessionに渡す
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	return session, err
}

// func (メソッドで使う構造体の略称 構造体の型) メソッド名(引数なし) (返り値 返り値の型) {処理内容}
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	 from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	//「sessionが存在するか」の判定
	if err != nil { //エラーがある場合
		valid = false //validをfalseにして
		return        //そのままreturnする
	}
	if sess.ID != 0 { //sess.IDが「0」でない場合
		valid = true //validをtrueにして
	}
	return valid, err //valid と err を返す
}

// UUIDと一致するsessionを削除するメソッド
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}
