package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

func (h *handler) Modify(rw http.ResponseWriter, r *http.Request) {
	u := new(User)

	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if u.Name == "" {
		http.Error(rw, ErrNameRequire.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, int64(id), u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO get id

	rw.Header().Add("Content-type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func Update(db *sql.DB, id int64, u *User) error {
	u.ModifiedAt = time.Now()

	stmt := `UPDATE "users" SET "name"=$1, "modified_at"=%2 WHERE id=%3)`

	_, err := db.Exec(stmt, u.Name, u.ModifiedAt, id)

	return err

}
