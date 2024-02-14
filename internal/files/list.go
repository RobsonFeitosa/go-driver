package files

import "database/sql"

func List(db *sql.DB, folderID int64) ([]File, error) {
	stmt := `SELECT * FROM "files" WHERE folder_id = $1`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	files := make([]File, 0)
	for rows.Next() {
		var f File

		err := rows.Scan(&f.ID, &f.FolderID, &f.OwnerID, &f.CreatedAt, &f.ModifiedAt, &f.Deleted)
		if err != nil {
			continue
		}

		files = append(files, f)
	}

	return files, nil
}