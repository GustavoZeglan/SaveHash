package password

import (
	"github.com/GustavoZeglan/SaveHash/core/password/domain"
	"strconv"
)

type UseCase interface {
	FindByUserId(userId string) ([]domain.Password, error)
	InsertPassword(password *domain.PasswordRequest, userId string) (int, error)
}

type PasswordService struct {
	Repo PasswordRepository
}

func NewService(repo PasswordRepository) PasswordService {
	return PasswordService{Repo: repo}
}

func (ps PasswordService) InsertPassword(password *domain.PasswordRequest, userId string) (int, error) {
	uId, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}

	p := &domain.Password{
		Name:   password.Name,
		Hash:   password.Hash,
		UserID: uId,
	}

	id, err := ps.Repo.Save(p)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ps *PasswordService) FindByUserId(userId string) ([]domain.Password, error) {

	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	passwords, err := ps.Repo.FindByUserId(id)
	if err != nil {
		return nil, err
	}

	return passwords, nil
}
