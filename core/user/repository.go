package user

import (
	"database/sql"
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
)

type UserRepository interface {
	Save(user *domain.User) (int, error)
	FindByEmail(email string) (domain.ResponseUser, error)
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

func (r *userRepository) Save(u *domain.User) (int, error) {
	query, err := r.DB.Prepare("INSERT INTO users(user_name, email, password) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return 0, err
	}

	defer query.Close()

	var id int
	err = query.QueryRow(u.Username, u.Email, u.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *userRepository) FindByEmail(email string) (domain.ResponseUser, error) {
	var u domain.ResponseUser
	query, err := r.DB.Prepare("SELECT id, user_name, email FROM users WHERE email = $1")
	if err != nil {
		return domain.ResponseUser{}, err
	}

	defer query.Close()

	err = query.QueryRow(email).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return domain.ResponseUser{}, err
	}

	return u, nil
}
