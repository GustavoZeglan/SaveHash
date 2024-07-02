package handler

import (
	"encoding/json"
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/core/user/domain"
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
	var req domain.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	u, err := domain.NewUser(req.Username, req.Email, req.Password)
	if err != nil {
		msg := utils.Message{Message: "Invalid body request", Status: http.StatusBadRequest, Data: err.Error()}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	resUser, err := uh.service.SignUp(u.Username, u.Email, u.Password)

	if err != nil && err.Error() == "email already registered" {
		msg := utils.Message{Message: "Email or password are invalid", Status: http.StatusBadRequest}
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
	var req domain.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	loginReq, err := domain.NewLoginRequest(req.Email, req.Password)
	if err != nil {
		msg := utils.Message{Message: "Invalid body request", Status: http.StatusBadRequest, Data: err.Error()}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	storageUser, err := uh.service.GetUserByEmail(req.Email)
	if err != nil {
		msg := utils.Message{Message: "Unregistered user", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	authenticated, err := uh.service.Login(loginReq.Email, loginReq.Password)
	if err != nil || !authenticated {
		msg := utils.Message{Message: "Invalid credentials", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	token, err := utils.CreateToken(strconv.Itoa(storageUser.ID), loginReq.Email)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	loginResp := domain.NewLoginResponse(token)

	utils.RespondWithJSON(w, http.StatusAccepted, loginResp)
}
