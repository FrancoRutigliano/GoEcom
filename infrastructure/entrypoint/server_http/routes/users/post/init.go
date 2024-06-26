package post

import (
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

type UserPost struct {
	Handler users.Handler
}

func Init(r fiber.Router, h users.Handler) {
	u := UserPost{}

	r.Post("/new", u.CreateUser)
}
