package get

import "github.com/gofiber/fiber/v2"

func (u *UserGet) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "All users"})
}
