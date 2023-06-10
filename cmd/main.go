package main

import (
	"links/config"
	"links/pkg/db"
	"log"
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

	db.Connection, err = db.InitPsqlDB(conf)
	if err != nil {
		log.Fatal(err)
	}

	//err = db.Connection.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

}
