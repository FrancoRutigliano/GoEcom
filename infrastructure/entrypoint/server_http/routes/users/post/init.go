package post

import "github.com/gofiber/fiber/v2"

type IUserPost interface {
	CreateUser(c *fiber.Ctx) error
}

type UserPost struct {
}

func Init(r fiber.Router, i IUserPost) {
	r.Post("/new", i.CreateUser)
}
