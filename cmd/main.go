package main

import (
	"GoEcom/infrastructure/db"
	server "GoEcom/infrastructure/entrypoint/server_http"
	"log"

	"github.com/go-sql-driver/mysql"
)

func init() {
	conn, err := db.NewSqlStorage(mysql.Config{
		User:   "fran",
		Passwd: "1234",
		DBName: "GoEcom",
		Addr:   "127.0.0.1:3306",
		Net:    "tcp",
	})
	if err != nil {
		log.Fatal(err)
	}

	db.InitStorage(conn)
}

func main() {
	s := server.NewServer()

	if err := s.Run(); err != nil {
		log.Fatal("error running server: ", err)
	}
}
