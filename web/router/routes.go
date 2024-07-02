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

	// Instance of user Repository
	userRepo := user.NewUserRepository(db)
	// Instance of User service
	userService := user.NewService(userRepo)
	// Instance of User Handler
	userHandler := handler.NewUserHandler(userService)

	// Instance of Password Repository
	passwordRepo := password.NewPasswordRepository(db)
	// Instance of Password Service
	passwordService := password.NewService(passwordRepo)
	// Instance of Password Handler
	passwordHandler := handler.NewPasswordHandler(passwordService)

	r.Group(func(r chi.Router) {
		r.Post("/signup", userHandler.SignUp)
		r.Post("/login", userHandler.Login)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/password", passwordHandler.CreatePassword)
		r.Get("/passwords", passwordHandler.GetPasswords)
	})
}
