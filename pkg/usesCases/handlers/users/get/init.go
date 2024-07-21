package get

import (
	users "GoEcom/infrastructure/data/repositoryUsers"
	usersDomain "GoEcom/pkg/domain"
)

type Iget interface {
	GetUsers() ([]usersDomain.CustomerInfo, error)
	GetUserByName(string) (usersDomain.CustomerInfo, error)
	GetUserById(string) (usersDomain.CustomerInfo, error)
}

type Get struct {
	UserRepository users.Repository
}
