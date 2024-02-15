package files

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
	f := File{
		ID:   1,
		Name: "Robson.jpg",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&f)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockListFiles(ts.mock)

	setMockUpdateFile(ts.mock, &f)

	ts.handler.Modify(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestUpdate() {
	f := File{
		ID:   1,
		Name: "testimg.jpg",
	}
	setMockUpdateFile(ts.mock, &f)

	err := Update(ts.conn, &File{Name: "testimg.jpg"}, 1)
	assert.NoError(ts.T(), err)
}

func setMockUpdateFile(mock sqlmock.Sqlmock, entity *File) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "files" SET "name"=$1, "modified_at"=$2, "deleted"=%3 WHERE id=%4`)).
		WithArgs(entity.Name, AnyTime{}, false, entity.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
