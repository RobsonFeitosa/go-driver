package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestCreate() {

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(ts.entity)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", &b)

	ts.entity.SetPassword(ts.entity.Password)

	ts.mock.ExpectExec(`INSERT INTO "users" ("name", "login", "password", "modified_at")*`).
		WithArgs(ts.entity.Name, ts.entity.Login, ts.entity.Password, ts.entity.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestInsert() {
	ts.mock.ExpectExec(`INSERT INTO "users" ("name", "login", "password", "modified_at")*`).
		WithArgs("Robson", "robson.gw@hotmail.com", ts.entity.Password, ts.entity.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}
