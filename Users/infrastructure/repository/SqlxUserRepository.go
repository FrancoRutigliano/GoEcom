package infraUserRepository

import (
	usersDomain "GoEcom/Users/pkg/domain"

	"github.com/jmoiron/sqlx"
)

type SqlxUserRepository struct {
	db *sqlx.DB
}

func NewSqlxUserRepository(db *sqlx.DB) *SqlxUserRepository {
	return &SqlxUserRepository{
		db: db,
	}
}

func (s *SqlxUserRepository) GetUsers() ([]usersDomain.UserResponse, error) {
	res := []usersDomain.UserResponse{}

	err := s.db.Get(&res, "SELECT username, email FROM users")
	if err != nil {
		return nil, err
	}

	return res, nil
}
