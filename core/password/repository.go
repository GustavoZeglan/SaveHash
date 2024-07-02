package password

import (
	"database/sql"
	"github.com/GustavoZeglan/SaveHash/core/password/domain"
)

type PasswordRepository interface {
	Save(password *domain.Password) (int, error)
	FindByUserId(id int) ([]domain.Password, error)
}

type passwordRepository struct {
	DB *sql.DB
}

func NewPasswordRepository(db *sql.DB) PasswordRepository {
	return &passwordRepository{
		DB: db,
	}
}

func (p *passwordRepository) Save(password *domain.Password) (int, error) {
	var id int

	query, err := p.DB.Prepare("INSERT INTO passwords(name, user_id, hash) VALUES ($1, $2, $3) RETURNING id;")
	if err != nil {
		return id, err
	}

	defer query.Close()

	err = query.QueryRow(password.Name, password.UserID, password.Hash).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (p *passwordRepository) FindByUserId(userId int) ([]domain.Password, error) {
	query, err := p.DB.Prepare("SELECT id, hash, name, user_id FROM passwords WHERE user_id = $1")
	if err != nil {
		return []domain.Password{}, nil
	}

	defer query.Close()

	rows, err := query.Query(userId)
	if err != nil {
		return []domain.Password{}, nil
	}

	var passwords []domain.Password

	for rows.Next() {
		password := domain.Password{}
		err = rows.Scan(&password.ID, &password.Hash, &password.Name, &password.UserID)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}
