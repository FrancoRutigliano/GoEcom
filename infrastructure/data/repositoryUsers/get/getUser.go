package get

import users "GoEcom/pkg/domain"

var (
	name     = "Franco"
	lastName = "Rutigliano"
	email    = "test@test.com"
)

func (UserGet) GetUsers() users.CustomerInfo {

	return users.CustomerInfo{
		FirstName: name,
		LastName:  lastName,
		Email:     email,
	}
}
