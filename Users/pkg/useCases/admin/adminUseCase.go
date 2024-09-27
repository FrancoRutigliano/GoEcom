package adminUseCase

import (
	infraUserRepository "GoEcom/Users/infrastructure/repository"
	usersDomain "GoEcom/Users/pkg/domain"
	"GoEcom/shared/infrastructure/data"
)

type AdminUseCase interface {
	GetUsers() ([]usersDomain.UserResponse, error)
}

type Admin struct {
	repository usersDomain.UserRepository
}

func NewAdmin() *Admin {
	db := data.GetConnection()
	defer db.Close()

	repo := infraUserRepository.NewSqlxUserRepository(db)

	return &Admin{repository: repo}
}

func (a *Admin) GetUsers() ([]usersDomain.UserResponse, error) {

	response, err := a.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return response, nil
}
