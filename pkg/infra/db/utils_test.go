package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	realTestValueUsername   = "username1"
	realTestValueHashedPass = "213PretendHashed"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return dbMock, mock
}

func newMockAddUser(t *testing.T, m sqlmock.Sqlmock) sqlmock.Sqlmock {

	// m.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// m.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mockQueryGetUserRows := sqlmock.NewRows(
		[]string{"id"},
	)
	mockExecNewUserRows := sqlmock.NewResult(0, 1)

	// regexp.QuoteMeta()
	m.ExpectQuery(fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1", userColumnID, userTableUsers, userColumnUsername)).WithArgs(
		realTestValueUsername,
	).WillReturnRows(mockQueryGetUserRows)
	m.ExpectExec(fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES($1, $2)", userTableUsers, userColumnUsername, userColumnPasswordHash)).
		WithArgs(
			realTestValueUsername,
			realTestValueHashedPass,
		).WillReturnResult(mockExecNewUserRows)

	return m
}
