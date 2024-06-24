package post

import (
	dom "GoEcom/pkg/domain"
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

func (u *UserPost) CreateUser(c *fiber.Ctx) error {
	s := dom.CreateUserRequest{}

	if err := c.BodyParser(&s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}

	r := users.Handler{}

	r.NewUserHandler()

	response := r.Post.CreateUser(s)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
