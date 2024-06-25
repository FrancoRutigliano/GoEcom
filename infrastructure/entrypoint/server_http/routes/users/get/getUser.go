package get

import (
	"github.com/gofiber/fiber/v2"
)

func (u *UserGet) GetUsers(c *fiber.Ctx) error {

	response := u.Handler.Get.GetUsers()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
