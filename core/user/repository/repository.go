package repository

import (
	"database/sql"
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
	"github.com/GustavoZeglan/SaveHash/core/user/entity"
	"github.com/GustavoZeglan/SaveHash/core/user/repository/converter"
)

type UserRepository interface {
	Save(user domain.UserDomainInterface) (domain.UserDomainInterface, error)
	FindByEmail(email string) (domain.UserDomainInterface, error)
	FindPassword(email string) (string, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindPassword(email string) (string, error) {
	query, err := r.DB.Prepare("SELECT password FROM users WHERE email = $1")
	if err != nil {
		return "", err
	}

	defer query.Close()

	var hashedPassword string

	err = query.QueryRow(email).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}

func (r *userRepository) Save(di domain.UserDomainInterface) (domain.UserDomainInterface, error) {
	query, err := r.DB.Prepare("INSERT INTO users(user_name, email, password) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return nil, err
	}

	u := converter.ConvertDomainToEntity(di)

	defer query.Close()

	err = query.QueryRow(u.Username, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(u), nil
}

func (r *userRepository) FindByEmail(email string) (domain.UserDomainInterface, error) {
	u := &entity.UserEntity{}
	query, err := r.DB.Prepare("SELECT id, user_name, email FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	err = query.QueryRow(email).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return nil, err
	}

	return converter.ConvertEntityToDomain(u), nil
}
