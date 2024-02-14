package files

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	entity, err := New(1, "robson", "jpg", "/path")
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at")*`).
		WithArgs(0, 1, "robson", "jpg", "/path", entity.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, entity)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
