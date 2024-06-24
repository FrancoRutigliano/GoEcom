package get

import (
	"GoEcom/pkg/usesCases/handlers/users"

	"github.com/gofiber/fiber/v2"
)

func (u *UserGet) GetUsers(c *fiber.Ctx) error {

	h := users.Handler{}
	h.NewUserHandler()
	response := h.Get.GetUsers()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": response})
}
