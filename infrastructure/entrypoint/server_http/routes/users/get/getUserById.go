package get

import "github.com/gofiber/fiber/v2"

func (u UserGet) GetUserById(c *fiber.Ctx) {
	response, err := u.Handler.Get.GetUserById(c.Params("id"))

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
