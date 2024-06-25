package users

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes/users/get"
	"GoEcom/infrastructure/entrypoint/server_http/routes/users/post"
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

func Init(r fiber.Router) {

	handlers := users.Handler{}
	handlers.NewUserHandler()

	get.Init(r, handlers)
	post.Init(r, handlers)
}
