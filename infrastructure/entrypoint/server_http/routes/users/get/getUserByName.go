package get

import (
	"github.com/gofiber/fiber/v2"
)

func (u *UserGet) GetUserByName(c *fiber.Ctx) error {
	response, err := u.Handler.Get.GetUserByName(c.Params("name"))

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
