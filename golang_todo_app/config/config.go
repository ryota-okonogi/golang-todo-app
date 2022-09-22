package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
	// "todo_app/todo_app/utils"
)

type ConfigList struct {
	Port      string //サーバーのポート番号
	SQLDriver string //SQLのソース
	DbName    string //データベース名
	LogFile   string //ログを残すファイル
}

var Config ConfigList //構造体の「ConfigList」を外部から呼び出せるようにする為、グローバルに変数宣言する。

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("golang_todo_app/config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
