package users

import (
	"GoEcom/pkg/usesCases/handlers/users/get"
	"GoEcom/pkg/usesCases/handlers/users/post"
)

type Handler struct {
	Get  get.Iget
	Post post.IUserPost
}

func (h *Handler) NewUserHandler() {
	h.Get = &get.Get{}
	h.Post = &post.Post{}
}
