package handler

import (
	"encoding/json"
	"fmt"
	"github.com/GustavoZeglan/SaveHash/core/password"
	"github.com/GustavoZeglan/SaveHash/web/model"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"net/http"
	"strconv"
)

type PasswordHandler struct {
	service *password.PasswordService
}

func NewPasswordHandler(service password.PasswordService) *PasswordHandler {
	return &PasswordHandler{&service}
}

func (ph *PasswordHandler) CreatePassword(w http.ResponseWriter, r *http.Request) {
	p, _ := r.Context().Value("payload").(model.PasswordRequest)

	payload := r.Header.Get("user_id")
	if payload == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization header not set"))
		return
	}

	userId, err := strconv.Atoi(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An unexpected error was occurred"))
		return
	}

	newPassword := password.NewPassword(p.Name, p.Hash, userId)

	id, err := ph.service.InsertPassword(newPassword)
	if err != nil {
		fmt.Println(newPassword)
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	newPassword.ID = uint64(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	res, err := json.Marshal(newPassword)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}
	w.Write(res)
}

func (ph *PasswordHandler) GetPasswords(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user_id")

	if userId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization header not set"))
		return
	}

	passwords, err := ph.service.FindByUserId(userId)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(passwords)
	w.Write(res)
}
