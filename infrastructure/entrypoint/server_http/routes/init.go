package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

}
