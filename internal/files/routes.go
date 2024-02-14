package files

import (
	"database/sql"

	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/RobsonFeitosa/go-driver/internal/queue"
	"github.com/go-chi/chi"
)

type handler struct {
	db     *sql.DB
	bucket *bucket.Bucket
	queue  *queue.Queue
}

func SetRoutes(r chi.Router, db *sql.DB, b *bucket.Bucket, q *queue.Queue) {
	h := handler{db, b, q}

	r.Put("/{id}", h.Modify)
}
