package files

import (
	"database/sql"
	"time"
)

func Update(db *sql.DB, f *File, id int64) error {
	f.ModifiedAt = time.Now()

	stmt := `UPDATE "files" SET "name"=$1, "modify_at"=$2, "deleted"=%3 WHERE id=%4)`

	_, err := db.Exec(stmt, f.Name, f.ModifiedAt,.Deleted, id)

	return err
}