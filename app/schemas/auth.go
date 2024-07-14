package schemas

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
