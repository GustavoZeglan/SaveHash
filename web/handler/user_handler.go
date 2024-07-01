package handler

import (
	"encoding/json"
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req user.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	u, err := user.NewUser(req.Username, req.Email, req.Password)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest, Data: err.Error()}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	resUser, err := uh.service.SignUp(u.Username, u.Email, u.Password)

	if err != nil && err.Error() == "email already registered" {
		msg := utils.Message{Message: "Probally that email address is already in use", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if err != nil {
		msg := utils.Message{Message: "Validation error", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, resUser)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req user.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	lr, err := user.NewLoginRequest(req.Email, req.Password)
	if err != nil {
		msg := utils.Message{Message: "Validation error", Status: http.StatusBadRequest, Data: err.Error()}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	storageUser, err := uh.service.GetUserByEmail(req.Email)
	if err != nil {
		msg := utils.Message{Message: "Unregistered user", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	authenticated, err := uh.service.Login(lr.Email, lr.Password)
	if err != nil || !authenticated {
		msg := utils.Message{Message: "Invalid credentials", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	token, err := utils.CreateToken(strconv.Itoa(storageUser.ID), lr.Email)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	resp := map[string]string{"token": token}

	msg := utils.Message{Message: "Successfully logged in", Status: http.StatusAccepted, Data: resp}
	utils.RespondWithJSON(w, http.StatusAccepted, msg)
}
