package folders

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type TransactionSuite struct {
	suite.Suite
	conn *sql.DB
	mock sqlmock.Sqlmock

	handler handler
	entity  *Folder
}

func (ts *TransactionSuite) SetupTest() {
	var err error

	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	ts.handler = handler{ts.conn}

	ts.entity = &Folder{
		Name: "Fotos",
	}

	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) AfterTest(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}

func setMockListFile(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "robson", "jpg", "/", time.Now(), time.Now(), false).
		AddRow(2, 1, 1, "ana", "jpg", "/", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "files" WHERE "folder_id" = $1 and "deleted"=false`)).
		WillReturnRows(rows)
}

func setMockGetFolder(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(1, 2, "Documentos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)
}

func setMockListFolders(mock sqlmock.Sqlmock) {
	foldersRows := sqlmock.NewRows([]string{"id", "parent_id", "name", "created_at", "modified_at", "deleted"}).
		AddRow(2, 3, "Projetos", time.Now(), time.Now(), false).
		AddRow(4, 3, "Trabalhos", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "folders" where parent_id=$1 and "deleted"=false`)).
		WithArgs(1).
		WillReturnRows(foldersRows)
}
