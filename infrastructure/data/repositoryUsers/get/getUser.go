package get

import (
	users "GoEcom/pkg/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

func (UserGet) GetUsers(db *sqlx.DB) ([]users.CustomerInfo, error) {
	clients := []users.CustomerInfo{}
	rows, err := db.Queryx("SELECT first_name, last_name, email FROM users ORDER BY first_name ASC")
	if err != nil {
		log.Fatal("error executing query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var client users.CustomerInfo
		err := rows.StructScan(&client)
		if err != nil {
			log.Fatal("error scanning struct: ", err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}
