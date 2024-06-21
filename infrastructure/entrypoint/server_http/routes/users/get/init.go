package get

import (
	"github.com/gofiber/fiber/v2"
)

type IUserGet interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}

// UserGet implementa IUserGet
type UserGet struct{}

func (u *UserGet) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "All users"})
}

func (u *UserGet) GetUserByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get user by ID"})
}

func Init(r fiber.Router, i IUserGet) {
	r.Get("/all", i.GetUsers)
	r.Get("/:id", i.GetUserByID)
}
