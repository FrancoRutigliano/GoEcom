package post

import dom "GoEcom/pkg/domain"

func (p *UserPost) Create(user dom.CreateUserRequest) dom.CreateUserResponse {
	return dom.CreateUserResponse{
		FirstName: user.FirstName,
		Email:     user.Email,
	}
}
