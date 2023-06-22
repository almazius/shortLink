package main

import (
	"google.golang.org/grpc"
	"links/config"
	handlers "links/internal/links/server"
	api "links/pkg/api/proto"
	"links/pkg/db"
	"log"
	"net"
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

	err = db.Connection.Ping()
	if err != nil {
		db.Log.Fatal(err)
	}

	//server := grpc.NewServer()
	//srv := &handlers.ShortLinkServer{}
	//api.RegisterShortLinkServer(server, srv)
	//l, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if server.Serve(l) != nil {
	//	log.Fatal(err)
	//}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := handlers.ShortLinkServer{}
	api.RegisterShortLinkServer(s, &srv)
	//RegisterGreeterServer(s, &api.ShortLinkServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
