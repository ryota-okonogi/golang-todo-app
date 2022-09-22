package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //logFileを読み込んで、オプションを指定(読み書き|作成|追記, パーミッション)して変数を作成。
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   //ログの書き込み先を指定(標準出力, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //ログのフォーマットを指定(日付 | 時刻 | ファイルの行番号)
	log.SetOutput(multiLogFile)                          //ログの出力先を指定
}
