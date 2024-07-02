package domain

import (
	"errors"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"strings"
)

type Password struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" validate:"required,max=50"`
	Hash   string `json:"hash" validate:"required,max=255"`
	UserID int    `json:"user_id" validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewPassword(name string, hash string, userID int) (*Password, error) {
	p := &Password{
		Name:   name,
		Hash:   hash,
		UserID: userID,
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Password) Validate() error {

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
