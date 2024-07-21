package adminUseCase

import (
	usersDomain "GoEcom/Users/pkg/domain"
	"GoEcom/shared/infrastructure/data"
)

type AdminUseCase interface {
	GetUsers() ([]usersDomain.UserResponse, error)
}

type Admin struct {
	//db *sqlx.DB
}

func (a *Admin) GetUsers() ([]usersDomain.UserResponse, error) {
	// conexion
	_ = data.GetConnection()
	// repositorio
	return []usersDomain.UserResponse{
		{Username: "admin", Email: "admin"},
	}, nil
}
