package user

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type UseCase interface {
	Login(email, password string) (string, error)
	SignUp(username, email, password string) error
	GetAllUsers() ([]User, error)
	GetUserByEmail(email string) (User, error)
}

type UserService struct {
	DB *sql.DB
}

func NewService(DB *sql.DB) *UserService {
	return &UserService{DB}
}

func (us *UserService) Login(email, password string) (bool, error) {

	query, err := us.DB.Prepare("SELECT password FROM users WHERE email = $1")
	if err != nil {
		return false, err
	}

	defer query.Close()

	var hashedPassword string

	err = query.QueryRow(email).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (us *UserService) SignUp(username, email, password string) (*ResponseUser, error) {

	storageUser, err := us.GetUserByEmail(email)
	if storageUser.Email == email {
		return nil, errors.New("email already registered")
	}

	u, _ := NewUser(username, email, password)

	query, err := us.DB.Prepare("INSERT INTO users(user_name, email, password) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var id int
	err = query.QueryRow(u.Username, u.Email, u.Password).Scan(&id)
	if err != nil {
		return nil, err
	}

	ru := NewResponseUser(id, u.Username, u.Email)

	return ru, nil
}

func (us *UserService) GetUserByEmail(email string) (ResponseUser, error) {
	var u ResponseUser
	query, err := us.DB.Prepare("SELECT id, user_name, email FROM users WHERE email = $1")
	if err != nil {
		return ResponseUser{}, err
	}

	defer query.Close()

	err = query.QueryRow(email).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return ResponseUser{}, err
	}

	return u, nil
}
