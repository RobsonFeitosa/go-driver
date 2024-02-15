package folders

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestList() {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/{id}", nil)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockFoldersParentIdNull(ts.mock)
	setMockFilesFolderIdNull(ts.mock)

	ts.handler.List(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)
}
func (ts *TransactionSuite) TestGetRootSubFolders() {
	setMockFoldersParentIdNull(ts.mock)

	_, err := getRootSubFolder(ts.conn)
	assert.NoError(ts.T(), err)
}

func setMockFoldersParentIdNull(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, "Documentos", time.Now(), time.Now(), false).
		AddRow(5, nil, "Imagens", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where "parent_id" is null and "deleted"=false`)).
		WillReturnRows(rows)
}

func setMockFilesFolderIdNull(mock sqlmock.Sqlmock) {
	filesRows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, 1, "robson", "jpg", "/", time.Now(), time.Now(), false).
		AddRow(2, nil, 1, "ana", "jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "files" WHERE  "folder_id" is null and "deleted"=false`)).
		WillReturnRows(filesRows)
}
