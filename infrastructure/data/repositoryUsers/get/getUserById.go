package get

import (
	usersDomain "GoEcom/pkg/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

func (UserGet) GetUserById(id string, db *sqlx.DB) (usersDomain.CustomerInfo, error) {
	client := usersDomain.CustomerInfo{}

	err := db.Get(&client, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal("error executing query: ", err)
	}

	return client, nil
}
