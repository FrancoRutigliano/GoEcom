package get

import (
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

// UserGet implementa IUserGet
type UserGet struct {
	Handler users.Handler
}

func Init(r fiber.Router, h users.Handler) {
	u := UserGet{}
	u.Handler = h
	r.Get("/all", u.GetUsers)
	r.Get("/:name", u.GetUserByName)
	r.Get("/:id")
}
