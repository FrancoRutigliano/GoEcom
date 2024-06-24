package post

import dom "GoEcom/pkg/domain"

type IUserPost interface {
	Create(user dom.CreateUserRequest) dom.CreateUserResponse
}

type UserPost struct {
}
