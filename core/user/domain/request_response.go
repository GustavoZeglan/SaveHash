package domain

import (
	"errors"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type ResponseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func init() {
	validate = validator.New()
}

func (lr *LoginRequest) Validate() error {
	var validationErrors []string

	if err := validate.Struct(lr); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ReqError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		return errors.New(strings.Join(validationErrors, ", "))
	}

	return nil
}

func (sr *SignUpRequest) Validate() error {
	var validationErrors []string

	if err := validate.Struct(sr); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el utils.ReqError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
		}
		return errors.New(strings.Join(validationErrors, ", "))
	}

	return nil
}

func NewLoginRequest(email, password string) (*LoginRequest, error) {
	lr := &LoginRequest{
		Email:    email,
		Password: password,
	}

	err := lr.Validate()
	if err != nil {
		return nil, err
	}

	return lr, nil
}

func NewResponseUser(id int, username string, email string) *ResponseUser {
	return &ResponseUser{
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
