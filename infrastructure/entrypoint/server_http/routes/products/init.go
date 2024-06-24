package products

import (
	"github.com/gofiber/fiber/v2"
)

type ProductsRoutes struct {
}

func NewProductsRoutes() *ProductsRoutes {
	return &ProductsRoutes{}
}

func Init(app fiber.Router) {
	_ = NewProductsRoutes()
}
