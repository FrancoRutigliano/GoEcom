package get

import (
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

// UserGet implementa IUserGet
type UserGet struct{}

func Init(r fiber.Router, h users.Handler) {
	u := UserGet{}

	r.Get("/all", u.GetUsers)
	r.Get("/:id", u.GetUserByID)
}
