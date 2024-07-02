package domain

import (
	"errors"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"strings"
)

type PasswordRequest struct {
	Name string `json:"name" validate:"required,max=50"`
	Hash string `json:"hash" validate:"required,max=255"`
}

type PasswordResponse struct {
	ID   int    `json:"id"`
	Name string `json:"password"`
	Hash string `json:"hash"`
}

func NewPasswordRequest(name string, hash string) (*PasswordRequest, error) {
	p := &PasswordRequest{
		Name: name,
		Hash: hash,
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *PasswordRequest) Validate() error {

	var validationErrors []string

	if err := validate.Struct(p); err != nil {
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

func NewPasswordResponse(id int, name string, hash string) *PasswordResponse {
	return &PasswordResponse{
		ID:   id,
		Name: name,
		Hash: hash,
	}
}
