package routes

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes/cart"
	"GoEcom/infrastructure/entrypoint/server_http/routes/orders"
	"GoEcom/infrastructure/entrypoint/server_http/routes/products"
	"GoEcom/infrastructure/entrypoint/server_http/routes/users"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Users
	users.Init(v1.Group("/users"))

	// Products
	products.Init(v1.Group("/products"))

	// Orders
	orders.Init(v1.Group("/orders"))

	// Cart
	cart.Init(v1.Group("/cart"))
}
