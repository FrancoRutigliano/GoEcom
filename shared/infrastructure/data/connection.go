package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=user dbname=mydb password=password sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	log.Println("connected to database")

	return db
}
