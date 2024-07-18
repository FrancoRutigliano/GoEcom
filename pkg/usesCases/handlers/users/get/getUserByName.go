package get

import (
	"GoEcom/infrastructure/data"
	users "GoEcom/pkg/domain"
	"strings"
)

func (g *Get) GetUserByName(n string) users.CustomerInfo {
	db := data.GetConnection()
	g.UserRepository.NewUserRepository()

	response, _ := g.UserRepository.Get.GetUserByName(strings.ToLower(n), db)
	return response
}
