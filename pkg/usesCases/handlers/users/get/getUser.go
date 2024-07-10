package get

import (
	"GoEcom/infrastructure/data"
	users "GoEcom/pkg/domain"
)

func (g *Get) GetUsers() []users.CustomerInfo {

	db := data.GetConnection()

	g.UserRepository.NewUserRepository()
	response, _ := g.UserRepository.Get.GetUsers(db)
	return response
}
