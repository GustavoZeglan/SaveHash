package handler

import (
	"encoding/json"
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"net/http"
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
	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	errors, err := utils.ErrorHandler(u)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errors)
		return
	}

	err = uh.service.SignUp(u)
	if err != nil {
		msg := utils.Message{Message: "Probally that email address is already in use", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	msg := utils.Message{Message: "User created successfully", Status: http.StatusCreated, Data: u}
	utils.RespondWithJSON(w, http.StatusCreated, msg)
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := uh.service.GetAllUsers()
	if err != nil {
		msg := utils.Message{Message: "Internal Server Error", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user user.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		msg := utils.Message{Message: "An error occurred", Status: http.StatusBadRequest}
		utils.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	errors, err := utils.ErrorHandler(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errors)
		return
	}

	isMatch, err := uh.service.Login(user.Email, user.Password)
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

	token, err := utils.CreateToken(user.Username, user.Email)
	if err != nil {
		msg := utils.Message{Message: "An unexpected error was occurred", Status: http.StatusInternalServerError}
		utils.RespondWithJSON(w, http.StatusInternalServerError, msg)
		return
	}

	resp := map[string]string{"token": token}

	msg := utils.Message{Message: "Successfully logged in", Status: http.StatusAccepted, Data: resp}
	utils.RespondWithJSON(w, http.StatusAccepted, msg)
}
