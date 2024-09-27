package usersDomain

type UserResponse struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type UserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
