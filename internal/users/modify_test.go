package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestModify() {
	u := User{
		ID:   1,
		Name: "Robson",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&u)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockUpdate(ts.mock, ts.entity)
	setMockGet(ts.mock, ts.entity)

	ts.handler.Modify(rr, req)

	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestUpdate() {
	setMockUpdate(ts.mock, ts.entity)

	err := Update(ts.conn, 1, &User{Name: "Robson"})
	assert.NoError(ts.T(), err)
}
