package users

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes/users/get"

	"github.com/gofiber/fiber/v2"
)

type UsersRoutes struct {
	Get get.IUserGet
}

func NewUserRoutes() *UsersRoutes {
	return &UsersRoutes{
		Get: &get.UserGet{},
	}
}

func Init(r fiber.Router) {

	userRoutes := NewUserRoutes()

	get.Init(r, userRoutes.Get)
}
