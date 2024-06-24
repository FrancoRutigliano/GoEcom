package get

import users "GoEcom/pkg/domain"

func (g *Get) GetUsers() users.CustomerInfo {
	g.UserRepository.NewUserRepository()
	response := g.UserRepository.Get.GetUsers()
	return response
}
