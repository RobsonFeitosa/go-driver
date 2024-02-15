package folders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
)

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	h := handler{db}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "robson", "jpg", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "ana", "jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "files" WHERE "folder_id" = $1 and "deleted"=false`)).
		WillReturnRows(rows)

	foldersRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, 3, "Projetos", time.Now(), time.Now(), false).
		AddRow(4, 3, "Trabalhos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where parent_id=$1 and "deleted"=false`)).
		WithArgs(1).
		WillReturnRows(foldersRows)

	mock.ExpectExec(regexp.QuoteMeta(`update "folders" set "modified_at"=$1, "deleted"=%2 where id=$3`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	h.Delete(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Error: %v", rr)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
	}
}
