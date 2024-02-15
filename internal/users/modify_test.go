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
	tcs := []struct {
		ID                string
		MockID            int64
		WithMock          bool
		WithUser          *User
		MockUpdateWithErr bool
		MockGetWithErr    bool
		ExpectStatusCode  int
	}{
		// success
		{"1", 1, true, &User{ID: 1, Name: "Robson"}, false, false, http.StatusOK},
		// error
		{"2", 2, false, nil, true, true, http.StatusBadRequest},
		{"10", 10, false, &User{ID: 10}, true, false, http.StatusBadRequest},
		{"A", 0, false, &User{Name: "Robson"}, true, false, http.StatusInternalServerError},
		// {"25", 25, true, &User{ID: 25, Name: "Robson"}, true, false, http.StatusInternalServerError},
		// {"500", true, &User{ID: 500, Name: "Robson"}, false, true, http.StatusInternalServerError},
	}

	for _, tc := range tcs {
		var b bytes.Buffer
		err := json.NewEncoder(&b).Encode(tc.WithUser)
		assert.NoError(ts.T(), err)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

		ctx := chi.NewRouteContext()
		ctx.URLParams.Add("id", tc.ID)

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

		if tc.WithMock {
			setMockUpdate(ts.mock, tc.WithUser, tc.MockID, tc.MockUpdateWithErr)
			setMockGet(ts.mock, tc.WithUser)
		}

		ts.handler.Modify(rr, req)

		assert.Equal(ts.T(), tc.ExpectStatusCode, rr.Code)
	}
}

func (ts *TransactionSuite) TestUpdate() {
	setMockUpdate(ts.mock, ts.entity, 1, false)

	err := Update(ts.conn, 1, &User{Name: "Robson"})
	assert.NoError(ts.T(), err)
}
