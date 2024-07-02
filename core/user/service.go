package user

import (
	"errors"
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type UseCase interface {
	Login(email, password string) (string, error)
	SignUp(username, email, password string) error
	GetUserByEmail(email string) (domain.ResponseUser, error)
}

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Login(email, password string) (bool, error) {

	hashedPassword, err := us.repo.FindPassword(email)
	if err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (us *UserService) SignUp(username, email, password string) (*domain.ResponseUser, error) {

	storageUser, err := us.repo.FindByEmail(email)
	if storageUser.Email == email {
		return nil, errors.New("email already registered")
	}

	u, err := domain.NewUser(username, email, password)
	if err != nil {
		return nil, err
	}

	id, err := us.repo.Save(u)
	if err != nil {
		return nil, err
	}

	ru := domain.NewResponseUser(id, u.Username, u.Email)

	return ru, nil
}

func (us *UserService) GetUserByEmail(email string) (*domain.ResponseUser, error) {

	u, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	respUser := domain.NewResponseUser(u.ID, u.Username, u.Email)

	return respUser, nil
}
