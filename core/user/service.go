package user

import (
	"database/sql"
	"errors"
	"fmt"
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

	hashedPassword := ""

	err = query.QueryRow(email).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (us *UserService) SignUp(user User) (int, error) {

	storageUser, err := us.GetUserByEmail(user.Email)
	if storageUser.Email == user.Email {
		return 0, errors.New("User already exists")
	}

	sb, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	u := &User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(sb),
	}

	query, err := us.DB.Prepare("INSERT INTO users(user_name, email, password) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer query.Close()

	var id int
	err = query.QueryRow(u.Username, u.Email, u.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (us *UserService) GetUserByEmail(email string) (User, error) {
	var u User
	query, err := us.DB.Prepare("SELECT id, user_name, email FROM users WHERE email = $1")
	if err != nil {
		return User{}, err
	}

	defer query.Close()

	err = query.QueryRow(email).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	return u, nil
}
