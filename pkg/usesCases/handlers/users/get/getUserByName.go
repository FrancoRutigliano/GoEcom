package get

import (
	"GoEcom/infrastructure/data"
	users "GoEcom/pkg/domain"
	"strings"
)

func (g *Get) GetUserByName(n string) (users.CustomerInfo, error) {
	db := data.GetConnection()
	//defer db.Close()
	g.UserRepository.NewUserRepository()

	response, err := g.UserRepository.Get.GetUserByName(strings.ToLower(n), db)
	return response, err
}
