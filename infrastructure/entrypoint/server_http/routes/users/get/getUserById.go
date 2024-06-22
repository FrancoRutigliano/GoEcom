package get

import "github.com/gofiber/fiber/v2"

func (u *UserGet) GetUserByID(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get user by ID"})
}
