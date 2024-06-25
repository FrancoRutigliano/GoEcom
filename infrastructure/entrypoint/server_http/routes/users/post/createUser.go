package post

import (
	dom "GoEcom/pkg/domain"

	"github.com/gofiber/fiber/v2"
)

func (u *UserPost) CreateUser(c *fiber.Ctx) error {
	s := dom.CreateUserRequest{}

	if err := c.BodyParser(&s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on request", "data": err})
	}

	response := u.Handler.Post.CreateUser(s)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
