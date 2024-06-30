package handler

import (
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/web/model"
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
	u, _ := r.Context().Value("payload").(user.User)

	id, err := uh.service.SignUp(u)
	if err != nil {
		msg := utils.Message{Message: "Probally that email address is already in use", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	res := model.ResponseUser{
		ID:       id,
		Username: u.Username,
		Email:    u.Email,
	}

	msg := utils.Message{Message: "User created successfully", Status: http.StatusCreated, Data: res}
	utils.RespondWithJSON(w, http.StatusCreated, msg)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	u, _ := r.Context().Value("payload").(user.User)

	storagedUser, err := uh.service.GetUserByEmail(u.Email)
	if err != nil {
		msg := utils.Message{Message: "Unregistered user", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	isMatch, err := uh.service.Login(u.Email, u.Password)
	if err != nil {
		msg := utils.Message{Message: "Invalid credentials", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if !isMatch {
		msg := utils.Message{Message: "Invalid credentials", Status: http.StatusUnauthorized}
		utils.RespondWithJSON(w, http.StatusUnauthorized, msg)
		return
	}

	token, err := utils.CreateToken(strconv.Itoa(storagedUser.ID), u.Email)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	resp := map[string]string{"token": token}

	msg := utils.Message{Message: "Successfully logged in", Status: http.StatusAccepted, Data: resp}
	utils.RespondWithJSON(w, http.StatusAccepted, msg)
}
