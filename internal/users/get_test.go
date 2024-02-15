package users

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestGetByID() {

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockGet(ts.mock, ts.entity)

	ts.handler.GetByID(rr, req)

	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestGet() {
	setMockGet(ts.mock, ts.entity)

	_, err := Get(ts.conn, 1)
	assert.NoError(ts.T(), err)
}
