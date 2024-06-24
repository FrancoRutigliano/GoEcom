package users

import "GoEcom/pkg/usesCases/handlers/users/get"

type Handler struct {
	Get get.Iget
}

func (h *Handler) NewUserHandler() {
	h.Get = &get.Get{}
}
