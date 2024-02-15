package folders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestDelete() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockListFile(ts.mock)
	setMockListFolders(ts.mock)
	setMockUpdateDelete(ts.mock)

	ts.handler.Delete(rr, req)
	assert.Equal(ts.T(), http.StatusNoContent, rr.Code)
}

func setMockUpdateDelete(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`update "folders" set "modified_at"=$1, "deleted"=%2 where id=$3`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
