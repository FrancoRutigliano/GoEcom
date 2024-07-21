package userRoutes

import (
	userController "GoEcom/Users/infrastructure/controller"
	adminUserController "GoEcom/Users/infrastructure/controller/admin"

	"github.com/gofiber/fiber/v2"
)

func Init(f fiber.Router) {
	a := adminUserController.AdminUserController{}
	a.NewAdminUserController()

	u := userController.UserController{}
	u.NewUserController()

	f.Get("/users", a.GetUsers)
}
