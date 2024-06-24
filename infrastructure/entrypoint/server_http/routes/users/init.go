package users

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes/users/get"
	"GoEcom/infrastructure/entrypoint/server_http/routes/users/post"

	"github.com/gofiber/fiber/v2"
)

type UsersRoutes struct {
	Get  get.IUserGet
	Post post.IUserPost
}

func NewUserRoutes() *UsersRoutes {
	return &UsersRoutes{
		Get:  &get.UserGet{},
		Post: &post.UserPost{},
	}
}

func Init(r fiber.Router) {

	userRoutes := NewUserRoutes()

	get.Init(r, userRoutes.Get)
	post.Init(r, userRoutes.Post)
}
