package post

import dom "GoEcom/pkg/domain"

func (p *Post) CreateUser(user dom.CreateUserRequest) dom.CreateUserResponse {
	p.Repository.NewUserRepository()
	response := p.Repository.Post.Create(user)
	return response
}
