package folders

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/RobsonFeitosa/go-driver/internal/files"
	"github.com/go-chi/chi"
)

func (h *handler) Delete(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = deleteFolderContent(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Delete(h.db, int64(id))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: list folders

	rw.Header().Add("Content-Type", "application/json")
}

func deleteFolderContent(db *sql.DB, folderID int64) error {
	err := deleteFiles(db, int64(folderID))
	if err != nil {
		return err
	}

	return deleteSubfolders(db, int64(folderID))
}

func deleteSubfolders(db *sql.DB, folderID int64) error {

	subfolders, err := getSubFolder(db, folderID)
	if err != nil {
		return err
	}

	removedFolders := make([]Folder, 0, len(subfolders))
	for _, sf := range subfolders {
		err := Delete(db, sf.ID)
		if err != nil {
			break
		}

		err = deleteFolderContent(db, sf.ID)
		if err != nil {
			Update(db, &sf, sf.ID)
			break
		}

		removedFolders = append(removedFolders, sf)
	}

	if len(subfolders) != len(removedFolders) {
		for _, rf := range removedFolders {
			Update(db, &rf, rf.ID)
		}
	}

	return nil
}

func deleteFiles(db *sql.DB, folderID int64) error {
	f, err := files.List(db, int64(folderID))
	if err != nil {
		return err
	}

	removedFiles := make([]files.File, 0, len(f))
	for _, file := range f {
		file.Deleted = true
		err := files.Update(db, &file, file.ID)
		if err != nil {
			break
		}

		removedFiles = append(removedFiles, file)
	}

	if len(f) != len(removedFiles) {
		for _, file := range removedFiles {
			file.Deleted = false
			files.Update(db, &file, file.ID)
		}

		return err
	}

	return nil
}

func Delete(db *sql.DB, id int64) error {
	stmt := `update "folders" set "modified_at"=$1, "deleted"=%2 where id=$3`
	_, err := db.Exec(stmt, time.Now(), id)

	return err

}
