package get

import (
	"github.com/gofiber/fiber/v2"
)

func (u *UserGet) GetUsers(c *fiber.Ctx) error {

	response, err := u.Handler.Get.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
