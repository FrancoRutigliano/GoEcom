package get

import users "GoEcom/pkg/domain"

type IUserGet interface {
	GetUsers() users.CustomerInfo
}

type UserGet struct {
}
