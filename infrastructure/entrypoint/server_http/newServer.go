package server

import (
	"GoEcom/infrastructure/entrypoint/server_http/routes"

	"github.com/gofiber/fiber/v2"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := fiber.New()

	routes.Init(app)

	return app.Listen(":8080")
}
