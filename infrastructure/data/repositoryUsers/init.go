package users

import (
	"GoEcom/infrastructure/data/repositoryUsers/delete"
	"GoEcom/infrastructure/data/repositoryUsers/get"
	"GoEcom/infrastructure/data/repositoryUsers/post"
	"GoEcom/infrastructure/data/repositoryUsers/put"
)

type Repository struct {
	Get    get.IUserGet
	Post   post.IUserPost
	Put    put.IUserPut
	Delete delete.IUserDelete
}

func (r *Repository) NewUserRepository() {
	r.Get = &get.UserGet{}
	r.Post = &post.UserPost{}
	r.Put = &put.UserPut{}
	r.Delete = &delete.UserDelete{}
}
