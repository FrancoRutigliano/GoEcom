package get

import (
	users "GoEcom/infrastructure/data/repositoryUsers"
	usersDomain "GoEcom/pkg/domain"
)

type Iget interface {
	GetUsers() []usersDomain.CustomerInfo
	GetUserByName(string) usersDomain.CustomerInfo
}

type Get struct {
	UserRepository users.Repository
}
