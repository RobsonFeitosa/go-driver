package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestCreate() {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	ts.entity.SetPassword(ts.entity.Password)

	setMockInsert(ts.mock, ts.entity)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestInsert() {
	setMockInsert(ts.mock, ts.entity)

	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}
