package folders

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestModify() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&ts.entity)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockUpdateFolder(ts.mock)
	setMockGetFolder(ts.mock)

	ts.handler.Modify(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestUpdate() {
	setMockUpdateFolder(ts.mock)

	err := Update(ts.conn, &Folder{Name: "Fotos"}, 1)
	assert.NoError(ts.T(), err)
}

func setMockUpdateFolder(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "folders" SET "name"=$1, "modified_at"=%2 WHERE id=%3`)).
		WithArgs("Fotos", AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
