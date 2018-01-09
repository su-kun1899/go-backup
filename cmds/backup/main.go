package main

import (
	"errors"
	"log"
	"flag"
)

func main() {
	var fatalErr error
	defer func () {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()
	var (
		dbpath = flag.String("db", "./backupdata", "データベースディレクトリへのパス")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー；コマンドを指定して下さい")
		return
	}
}
