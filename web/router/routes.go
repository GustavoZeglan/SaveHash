package router

import (
	"database/sql"
	"github.com/GustavoZeglan/SaveHash/core/password"
	"github.com/GustavoZeglan/SaveHash/core/user"
	"github.com/GustavoZeglan/SaveHash/web/handler"
	"github.com/GustavoZeglan/SaveHash/web/middleware"
	"github.com/go-chi/chi/v5"
)

func initializeRoutes(r *chi.Mux, db *sql.DB) {

	// Instance of User service
	userService := user.NewService(db)
	// Instance of User Handler
	userHandler := handler.NewUserHandler(userService)

	// Instance of Password Service
	passwordService := password.NewService(db)

	// Instance of Password Handler
	passwordHandler := handler.NewPasswordHandler(passwordService)

	r.Post("/signup", middleware.ValidateBody[user.User](userHandler.SignUp))
	r.Post("/login", middleware.ValidateBody[user.User](userHandler.Login))

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/password", middleware.ValidateBody[password.Password](passwordHandler.CreatePassword))
		r.Get("/passwords", passwordHandler.GetPasswords)
	})
}
