package router

import (
	"database/sql"
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/web/handler"
	"github.com/GustavoZeglan/SaveHash/web/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func initializeRoutes(r *chi.Mux, db *sql.DB) {

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		msg := utils.Message{Message: "pong", Status: http.StatusOK}
		utils.RespondWithJSON(w, http.StatusOK, msg)
	})

	// User service
	userService := user.NewService(db)
	// User Handler
	userHandler := handler.NewUserHandler(userService)

	r.Get("/users", userHandler.GetUsers)
	r.Post("/signup", userHandler.SignUp)
	r.Post("/login", userHandler.Login)
}
