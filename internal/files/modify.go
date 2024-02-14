package files

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := Get(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = file.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = Update(h.db, file, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Add("Content-type", "application/json")
	json.NewEncoder(rw).Encode(file)
}

func Update(db *sql.DB, f *File, id int64) error {
	f.ModifiedAt = time.Now()

	stmt := `UPDATE "files" SET "name"=$1, "modify_at"=$2, "deleted"=%3 WHERE id=%4)`

	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, f.Deleted, id)

	return err
}
