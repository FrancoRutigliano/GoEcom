package routes

import (
	userRoutes "GoEcom/Users/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func Init(f *fiber.App) {
	api := f.Group("/api")
	userRoutes.Init(api)

}
