package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Initialize(db *sql.DB) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	initializeRoutes(r, db)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
