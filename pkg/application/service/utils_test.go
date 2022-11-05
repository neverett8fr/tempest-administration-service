package service

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"tempest-administration-service/pkg/infra/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

const (
	realTestValueUsername = "username1"
	realTestValuePass     = "123Password"

	userColumnID           = "id"
	userColumnUsername     = "username"
	userColumnPasswordHash = "password_hash"

	userTableUsers = "users"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return dbMock, mock
}

type anyPass struct{}

func (a anyPass) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func newMockAddUser(t *testing.T, m sqlmock.Sqlmock) sqlmock.Sqlmock {

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
			anyPass{},
		).WillReturnResult(mockExecNewUserRows)

	return m
}

func newMockRouter(t *testing.T, dbMock *sql.DB) *mux.Router {

	mockDBConn := db.NewDBConnFromExisting(dbMock)

	router := mux.NewRouter()
	NewServiceRoutes(router, mockDBConn.Conn)

	return router
}
