package response

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewResponseUser(id int, username string, email string) *UserResponse {
	return &UserResponse{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		Token: token,
	}
}
