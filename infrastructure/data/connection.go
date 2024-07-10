package data

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func GetConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://user:1234@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	return db
}
