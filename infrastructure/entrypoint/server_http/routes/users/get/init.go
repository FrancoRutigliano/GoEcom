package get

import "github.com/gofiber/fiber/v2"

type IUserGet interface {
	GetUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}

// UserGet implementa IUserGet
type UserGet struct{}

func (u *UserGet) GetUsers(c *fiber.Ctx) error {
	panic("implement me")
}

func (u *UserGet) GetUserByID(c *fiber.Ctx) error {
	panic("implement me")
}
