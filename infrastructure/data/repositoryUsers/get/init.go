package get

import (
	users "GoEcom/pkg/domain"

	"github.com/jmoiron/sqlx"
)

type IUserGet interface {
	GetUsers(*sqlx.DB) ([]users.CustomerInfo, error)
	GetUserByName(string, *sqlx.DB) (users.CustomerInfo, error)
	GetUserById(string, *sqlx.DB) (users.CustomerInfo, error)
}

type UserGet struct {
}
