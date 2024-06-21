package routes

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes/users"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	users.Init(v1.Group("/users"))
}
