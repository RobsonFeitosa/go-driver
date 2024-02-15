package users

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
	entity  *User
}

func (ts *TransactionSuite) SetupTest() {
	var err error

	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	ts.handler = handler{ts.conn}

	ts.entity = &User{
		Name:     "Robson",
		Login:    "robson.gw@hotmail.com",
		Password: "123456",
	}

	assert.NoError(ts.T(), err)
}

func (ts *TransactionSuite) AfterTest(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}

func setMockInsert(mock sqlmock.Sqlmock, entity *User) {
	mock.ExpectExec(`INSERT INTO "users" ("name", "login", "password", "modified_at")*`).
		WithArgs(entity.Name, entity.Login, entity.Password, entity.ModifiedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
}

func setMockGet(mock sqlmock.Sqlmock, entity *User) {
	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Robson", "robson.gw@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id=$1`)).
		WithArgs(1).
		WillReturnRows(rows)

	// if err {
	// 	exp.WillReturnError(sql.ErrNoRows)
	// } else {
	// 	exp.WillReturnRows(rows)
	// }
}

// rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
// AddRow(1, "Robson", "robson.gw@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now())

// mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id=$1`)).
// WithArgs(1).
// WillReturnRows(rows)
// }

func setMockUpdate(mock sqlmock.Sqlmock, entity *User, id int64, err bool) {
	exp := mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "name"=$1, "modified_at"=%2 WHERE id=%3`)).
		WithArgs(entity.Name, AnyTime{}, id)

	if err {
		exp.WillReturnError(sql.ErrNoRows)
	} else {
		exp.WillReturnResult(sqlmock.NewResult(1, 1))
	}
}

func setMockList(mock sqlmock.Sqlmock, entity *User) {
	rows := sqlmock.NewRows([]string{"id", "name", "login", "password", "created_at", "modified_at", "deleted", "last_login"}).
		AddRow(1, "Robson", "robson.gw@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now()).
		AddRow(2, "Ana", "ana@hotmail.com", "123456", time.Now(), time.Now(), false, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE deleted = false`)).
		WillReturnRows(rows)
}
