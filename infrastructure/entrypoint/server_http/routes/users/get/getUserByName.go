package get

import (
	"github.com/gofiber/fiber/v2"
)

func (u *UserGet) GetUserByName(c *fiber.Ctx) error {
	response := u.Handler.Get.GetUserByName(c.Params("name"))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
