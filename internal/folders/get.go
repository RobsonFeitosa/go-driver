package folders

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	f, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: get content
}

func Get(db *sql.DB, folderID int64) (*Folder, error) {
	stmt := `select * from "folders" where id=$1`
	row := db.QueryRow(stmt, folderID)

	var f Folder
	err := row.Scan(&f.ID, &f.ParentId, &f.Name, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
