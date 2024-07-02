package handler

import (
	"encoding/json"
	"github.com/GustavoZeglan/SaveHash/core/password"
	"github.com/GustavoZeglan/SaveHash/core/password/domain"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"net/http"
)

type PasswordHandler struct {
	service *password.PasswordService
}

func NewPasswordHandler(service password.PasswordService) *PasswordHandler {
	return &PasswordHandler{&service}
}

func (ph *PasswordHandler) CreatePassword(w http.ResponseWriter, r *http.Request) {
	payload := r.Header.Get("user_id")
	if payload == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization header not set"))
		return
	}

	var req domain.PasswordRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, nil)
		return
	}

	passwordRequest, err := domain.NewPasswordRequest(req.Name, req.Hash)
	if err != nil {
		msg := utils.Message{Message: "Invalid body request", Status: http.StatusBadRequest, Data: err.Error()}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	id, err := ph.service.InsertPassword(passwordRequest, payload)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	passResp := domain.NewPasswordResponse(id, passwordRequest.Name, passwordRequest.Hash)

	msg := utils.Message{Message: "Password successfully created", Status: http.StatusOK, Data: passResp}
	utils.RespondWithJSON(w, http.StatusOK, msg)
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
