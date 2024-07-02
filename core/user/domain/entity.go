package domain

import (
	"errors"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func NewUser(username, email, password string) (*User, error) {
	u := &User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := u.validate()
	if err != nil {
		return nil, err
	}

	err = u.Prepare()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Prepare() error {
	sb, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(sb)

	return nil
}

func (u *User) validate() error {

	var validationErrors []string

	if err := validate.Struct(u); err != nil {
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
