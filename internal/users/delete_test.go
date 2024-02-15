package users

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestDeleteHTTP() {
	tcs := []struct {
		ID               string
		WithMock         bool
		MockID           int64
		MockWithErr      bool
		ExpectStatusCode int
	}{
		{"1", true, 1, false, http.StatusNoContent},
		{"A", false, -1, true, http.StatusInternalServerError},
		{"25", true, 25, true, http.StatusInternalServerError},
	}

	for _, tc := range tcs {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/{id}", nil)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", tc.ID)

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		if tc.WithMock {
			setMockUpdateDelete(ts.mock, tc.MockID, tc.MockWithErr)
		}

		ts.handler.Delete(rr, req)
		assert.Equal(ts.T(), tc.ExpectStatusCode, rr.Code)
	}
}

func (ts *TransactionSuite) TestDelete() {
	setMockUpdateDelete(ts.mock, 1, false)

	err := Delete(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func setMockUpdateDelete(mock sqlmock.Sqlmock, id int64, err bool) {
	exp := mock.ExpectExec(`UPDATE "users" SET *`).
		WithArgs(AnyTime{}, id)

	if err {
		exp.WillReturnError(sql.ErrNoRows)
	} else {

		exp.WillReturnResult(sqlmock.NewResult(1, 1))
	}
}
