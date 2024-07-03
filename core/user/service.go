package user

import (
	"errors"
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
	"github.com/GustavoZeglan/SaveHash/core/user/repository"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type UseCase interface {
	Login(email, password string) (bool, error)
	SignUp(username, email, password string) (domain.UserDomainInterface, error)
	GetUserByEmail(email string) (domain.UserDomainInterface, error)
}

type UserService struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *UserService {
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

func (us *UserService) SignUp(username, email, password string) (domain.UserDomainInterface, error) {

	storageUser, err := us.repo.FindByEmail(email)
	if storageUser != nil && storageUser.GetEmail() == email {
		return nil, errors.New("email already registered")
	}

	ud := domain.NewUserDomain(0, username, email, password)
	ud.EncryptPassword()

	d, err := us.repo.Save(ud)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (us *UserService) GetUserByEmail(email string) (domain.UserDomainInterface, error) {

	ud, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return ud, nil
}
