package serverhttp

import (
	"GoEcom/shared/infrastructure/entrypoint/server_http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := fiber.New()

	app.Use(logger.New())

	// routes
	routes.Init(app)

	return app.Listen(":8080")
}
