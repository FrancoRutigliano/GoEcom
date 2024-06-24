package post

import (
	users "GoEcom/infrastructure/data/repositoryUsers"
	dom "GoEcom/pkg/domain"
)

type IUserPost interface {
	CreateUser(user dom.CreateUserRequest) dom.CreateUserResponse
}

type Post struct {
	Repository users.Repository
}
