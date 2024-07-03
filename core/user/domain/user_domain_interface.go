package domain

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetUsername() string
	GetId() int

	//SetID(string)

	EncryptPassword() error
	//GenerateToken() (string, error)
}

func NewUserDomain(id int, username, email, password string) UserDomainInterface {
	return &userDomain{
		id:       id,
		username: username,
		email:    email,
		password: password,
	}
}

func NewLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
