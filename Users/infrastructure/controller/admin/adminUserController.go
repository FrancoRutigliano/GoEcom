package adminUserController

import (
	adminUseCase "GoEcom/Users/pkg/useCases/admin"

	"github.com/gofiber/fiber/v2"
)

type AdminUserController struct {
	// use cases de admin
	adminUseCase adminUseCase.Admin
}

func (a *AdminUserController) NewAdminUserController() {
}

func (a *AdminUserController) GetUsers(f *fiber.Ctx) error {

	// interaccion con caso de uso admin
	response, err := a.adminUseCase.GetUsers()
	if err != nil {
		return f.Status(500).JSON(fiber.Map{"message": "oops something went wrong"})
	}
	if response == nil {
		return f.Status(404).JSON(fiber.Map{"message": "Not found"})
	}

	return f.Status(200).JSON(fiber.Map{"message": response})
}
