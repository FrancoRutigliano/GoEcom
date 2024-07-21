package get

import (
	"GoEcom/infrastructure/data"
	users "GoEcom/pkg/domain"
)

func (g *Get) GetUsers() ([]users.CustomerInfo, error) {

	db := data.GetConnection()

	g.UserRepository.NewUserRepository()
	response, err := g.UserRepository.Get.GetUsers(db)
	return response, err
}
