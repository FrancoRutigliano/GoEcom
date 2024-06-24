package users

type CustomerInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserResponse struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}
