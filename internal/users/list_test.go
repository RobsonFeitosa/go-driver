package users

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestList() {

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	setMockList(ts.mock, ts.entity)

	ts.handler.List(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}

func (ts *TransactionSuite) TestSelectAll() {
	setMockList(ts.mock, ts.entity)

	_, err := SelectAll(ts.conn)
	assert.NoError(ts.T(), err)
}
