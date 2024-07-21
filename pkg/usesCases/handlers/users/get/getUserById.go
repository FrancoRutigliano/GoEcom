package get

import (
	"GoEcom/infrastructure/data"
	usersDomain "GoEcom/pkg/domain"
)

func (g Get) GetUserById(id string) (usersDomain.CustomerInfo, error) {

	db := data.GetConnection()
	//defer db.Close()

	g.UserRepository.NewUserRepository()
	response, err := g.UserRepository.Get.GetUserById(id, db)

	return response, err
}
