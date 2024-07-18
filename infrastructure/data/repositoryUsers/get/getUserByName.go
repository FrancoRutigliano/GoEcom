package get

import (
	users "GoEcom/pkg/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

func (u *UserGet) GetUserByName(n string, db *sqlx.DB) (users.CustomerInfo, error) {
	client := users.CustomerInfo{}

	err := db.Get(&client, "SELECT first_name, last_name, email FROM users WHERE first_name = $1", n)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	return client, nil
}
