package folders

import (
	"database/sql"

	"github.com/RobsonFeitosa/go-driver/internal/auth"
	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

func SetRoutes(r chi.Router, db *sql.DB) {
	h := handler{db}

	r.Route("/folders", func(r chi.Router) {
		r.Use(auth.Validate)

		r.Post("/", h.Create)
		r.Put("/{id}", h.Modify)
		r.Delete("/{id}", h.Delete)
		r.Get("/{id}", h.Get)
		r.Get("/", h.List)
	})
}
