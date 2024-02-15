package folders

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestGet() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockGetFolder(ts.mock)
	setMockListFolders(ts.mock)
	setMockListFile(ts.mock)

	ts.handler.Get(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestGetFolder() {
	setMockGetFolder(ts.mock)

	_, err := GetFolder(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TestGetSubFolder() {
	setMockListFolders(ts.mock)

	_, err := getSubFolder(ts.conn, 1)
	assert.NoError(ts.T(), err)
}
