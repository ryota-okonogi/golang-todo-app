package config

import (
	"log"

	"golang-todo-app/golang_todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string //サーバーのポート番号
	SQLDriver string //SQLのソース
	DbName    string //データベース名
	LogFile   string //ログを残すファイル
	Static    string //静的ファイルがある階層をconfigに設定
}

var Config ConfigList //構造体の「ConfigList」を外部から呼び出せるようにする為、グローバルに変数宣言する。

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
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
		Static:    cfg.Section("web").Key("static").String(),
	}
}
