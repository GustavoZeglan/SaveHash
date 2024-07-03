package request

import (
	"github.com/go-playground/validator"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewSignupRequest(username, email, password string) (SignUpRequest, error) {

	v := Validator[SignUpRequest]{}

	sr := SignUpRequest{
		Username: username,
		Email:    email,
		Password: password,
	}

	v.body = sr

	err := v.Validate()
	if err != nil {
		return SignUpRequest{}, err
	}

	return sr, nil
}

func NewLoginRequest(email, password string) (LoginRequest, error) {

	v := Validator[LoginRequest]{}

	lr := LoginRequest{
		Email:    email,
		Password: password,
	}

	v.body = lr

	err := v.Validate()
	if err != nil {
		return LoginRequest{}, err
	}

	return lr, nil
}
