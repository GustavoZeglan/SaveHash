package password

type Password struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name" validate:"required,max=50"`
	Hash   string `json:"hash" validate:"required,max=255"`
	UserID int    `json:"user_id" validate:"required"`
}

func NewPassword(name string, hash string, userID int) *Password {
	return &Password{
		Name:   name,
		Hash:   hash,
		UserID: userID,
	}
}
