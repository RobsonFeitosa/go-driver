package files

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RobsonFeitosa/go-driver/internal/queue"
	"gopkg.in/guregu/null.v4"
)

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	path := fmt.Sprintf("/%s", fileHeader.Filename)

	fmt.Println("entrousaafa1")
	err = h.bucket.Upload(file, path)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("entrousaafa2")

	userID := r.Context().Value("user_id").(int64)

	entity, err := New(userID, fileHeader.Filename, fileHeader.Header.Get("Content-Type"), path)
	if err != nil {
		h.bucket.Delete(path)

		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	folderID := r.Form.Get("folder_id")
	if folderID != "" {
		fid, err := strconv.Atoi(folderID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		entity.FolderID = null.IntFrom(int64(fid))
	}

	id, err := Insert(h.db, entity)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	entity.ID = id

	dto := queue.QueueDto{
		Filename: fileHeader.Filename,
		Path:     path,
		ID:       int(id),
	}

	msg, err := dto.Marshal()
	if err != nil {
		// TODO: rollback
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.queue.Publish(msg)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(entity)
}

func Insert(db *sql.DB, f *File) (id int64, err error) {
	stmt := `insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = db.QueryRow(stmt, f.FolderID, f.OwnerID, f.Name,
		f.Type, f.Path, f.ModifiedAt).Scan(&id)
	if err != nil {
		return -1, err
	}

	return
}
