package folders

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	f := new(Folder)

	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = f.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, f, int64(id))

	// TODO get id

	rw.Header().Add("Content-type", "application/json")
	json.NewEncoder(rw).Encode(f)
}

func Update(db *sql.DB, f *Folder, id int64) error {
	f.ModifiedAt = time.Now()

	stmt := `UPDATE "folders" SET "name"=$1, "modified_at"=%2 WHERE id=%3)`

	_, err := db.Exec(stmt, f.Name, f.ModifiedAt, id)

	return err

}
