package model

type PasswordRequest struct {
	Name string `json:"name" validate:"required,max=50"`
	Hash string `json:"hash" validate:"required,max=255"`
}
