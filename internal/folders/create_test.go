package folders

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db}

	f := Folder{
		Name: "Fotos",
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(&f)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "folders" ("parent_id", "name", "modified_at") VALUES ($1, $2, $3)`)).
		WithArgs(0, "Fotos", AnyTime{}).
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

	f, err := New("Fotos", 0)
	if err != nil {
		t.Error(err)
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "folders" ("parent_id", "name", "modified_at") VALUES ($1, $2, $3)`)).
		WithArgs(0, "Fotos", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = Insert(db, f)
	if err != nil {
		t.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
