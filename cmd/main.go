package main

import (
	"fmt"
	"links/config"
	"links/internal/links/usecase"
	"links/pkg/db"
	"log"
	"os"
)

func main() {
	viperConf, err := config.LoadConfig() // загружаем конфиг для бд из папки config
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(viperConf)
	if err != nil {
		log.Fatal(err)
	}

	db.Log = log.New(os.Stdout, "BLAT ", log.Lshortfile|log.LstdFlags)

	db.Connection, err = db.InitPsqlDB(conf)
	if err != nil {
		db.Log.Fatal(err)
	}
	db.Connection.Exec(`drop table links`)
	db.InitTable()

	err = db.Connection.Ping()
	if err != nil {
		db.Log.Fatal(err)
	}

	s, err := usecase.POST("https://rocketchat-student.21-school.ru/direct/6CT56ZFNvQyT6DnKWATxdocjygq4DgPc6j")
	fmt.Println(s, err)
	fmt.Println(usecase.GET(s))

}
