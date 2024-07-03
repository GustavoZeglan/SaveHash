package domain

import "golang.org/x/crypto/bcrypt"

type userDomain struct {
	id       int
	username string
	email    string
	password string
}

func (ud *userDomain) GetId() int {
	return ud.id
}

func (ud *userDomain) GetUsername() string {
	return ud.username
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) EncryptPassword() error {
	sb, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	ud.password = string(sb)

	return nil
}
