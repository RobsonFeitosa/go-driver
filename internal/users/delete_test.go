package users

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestDeleteHTTP() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockUpdateDelete(ts.mock, ts.entity)

	ts.handler.Delete(rr, req)
	assert.Equal(ts.T(), http.StatusNoContent, rr.Code)

}

func (ts *TransactionSuite) TestDelete() {
	setMockUpdateDelete(ts.mock, ts.entity)

	err := Delete(ts.conn, 1)
	assert.NoError(ts.T(), err)
}
