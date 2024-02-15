package files

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RobsonFeitosa/go-driver/internal/bucket"
	"github.com/RobsonFeitosa/go-driver/internal/queue"
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
	entity  *File
}

func (ts *TransactionSuite) SetupTest() {
	var err error

	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	b, err := bucket.New(bucket.MockProvider, nil)
	assert.NoError(ts.T(), err)

	q, err := queue.New(queue.Mock, nil)
	assert.NoError(ts.T(), err)

	ts.handler = handler{ts.conn, b, q}

	ts.entity = &File{
		OwnerID: 1,
		Name:    "testimg.jpg",
		Type:    "application/octet-stream",
		Path:    "/testimg.jpg",
	}

	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) AfterTest(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}

func setMockListFiles(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"id", "folder_id", "owner_id", "name", "type", "path", "created_at", "modified_at", "deleted"}).
		AddRow(1, 1, 1, "robson", "jpg", "/src", time.Now(), time.Now(), false)

	mock.ExpectQuery(regexp.QuoteMeta(`select * from "files" where id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)
}
