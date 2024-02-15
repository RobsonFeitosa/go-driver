package users

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Robson", "robson.gw@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "Ana", "ana@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE deleted = false`)).
		WillReturnRows(rows)

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}

func TestSelectAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Robson", "robson.gw@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "Ana", "ana@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE deleted = false`)).
		WillReturnRows(rows)

	_, err = SelectAll(db)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
