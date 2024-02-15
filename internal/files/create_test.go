package files

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestCreate() {
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open("./testedata/testimg.jpg")
	assert.NoError(ts.T(), err)

	w, err := mw.CreateFormFile("file", "testimg.jpg")
	assert.NoError(ts.T(), err)

	_, err = io.Copy(w, file)
	assert.NoError(ts.T(), err)

	mw.Close()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	setMockInsert(ts.mock, ts.entity)

	ts.handler.Create(rr, req)
	assert.Equal(ts.T(), http.StatusCreated, rr.Code)
}

func (ts *TransactionSuite) TestInsert() {
	setMockInsert(ts.mock, ts.entity)

	_, err := Insert(ts.conn, ts.entity)
	assert.NoError(ts.T(), err)
}

func setMockInsert(mock sqlmock.Sqlmock, entity *File) {
	mock.ExpectExec(regexp.QuoteMeta(`insert into "files" ("folder_id", "owner_id", "name", "type", "path", "modified_at") values ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(0, 1, "testimg.jpg", "application/octet-stream", "/testimg.jpg", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
