package password

import (
	"database/sql"
)

type UseCase interface {
	FindByUserId(userId string) (string, error)
	InsertPassword(password Password) (int, error)
}

type PasswordService struct {
	DB *sql.DB
}

func NewService(db *sql.DB) PasswordService {
	return PasswordService{DB: db}
}

func (ps PasswordService) InsertPassword(password *Password) (int, error) {
	var id int

	query, err := ps.DB.Prepare("INSERT INTO passwords(name, user_id, hash) VALUES ($1, $2, $3) RETURNING id;")
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

func (ps *PasswordService) FindByUserId(userId string) ([]Password, error) {
	query, err := ps.DB.Prepare("SELECT id, hash, name, user_id FROM passwords WHERE user_id = $1")
	if err != nil {
		return []Password{}, nil
	}

	defer query.Close()

	rows, err := query.Query(userId)
	if err != nil {
		return []Password{}, nil
	}

	var passwords []Password

	for rows.Next() {
		password := Password{}
		err = rows.Scan(&password.ID, &password.Hash, &password.Name, &password.UserID)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, password)
	}

	return passwords, nil
}
