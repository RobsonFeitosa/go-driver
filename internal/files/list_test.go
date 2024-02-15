package files

import (
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func (ts *TransactionSuite) TestList() {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "robson", "jpg", "/src", time.Now(), time.Now(), false).
		AddRow(2, 2, 2, "ana", "jpg", "/src", time.Now(), time.Now(), false)

	ts.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "files" WHERE "folder_id" = $1 and "deleted"=false`)).
		WillReturnRows(rows)

	_, err := List(ts.conn, 1)
	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) TestListRoot() {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, nil, 1, "gopher", "image/jpg", "/", time.Now(), time.Now(), false).
		AddRow(2, nil, 2, "logo", "image/jpg", "/", time.Now(), time.Now(), false)

	ts.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "files" WHERE  "folder_id" is null and "deleted"=false`)).
		WillReturnRows(rows)

	_, err := ListRoot(ts.conn)
	assert.NoError(ts.T(), err)
}
