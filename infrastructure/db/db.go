package db

import (
	"database/sql"
	"log"
)

func InitStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
