package usersDomain

type UserResponse struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
