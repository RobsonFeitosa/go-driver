package files

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/RobsonFeitosa/go-driver/internal/queue"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	b, err := bucket.New(bucket.MockProvider, nil)
	if err != nil {
		t.Error(err)
	}

	q, err := queue.New(queue.Mock, nil)
	if err != nil {
		t.Error(err)
	}

	h := handler{db, b, q}

	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open("./testedata/testimg.jpg")
	if err != nil {
		t.Error(err)
	}

	w, err := mw.CreateFormFile("file", "testimg.jpg")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(w, file)
	if err != nil {
		t.Error(err)
	}

	mw.Close()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	mock.ExpectExec(regexp.QuoteMeta(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") values ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(0, 1, "testimg.jpg", "application/octet-stream", "/testimg.jpg", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

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
