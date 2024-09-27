package usersDomain

type UserRepository interface {
	GetUsers() ([]UserResponse, error)
	//GetUserById(db *sqlx.DB, id string) (UserResponse, error)
}
