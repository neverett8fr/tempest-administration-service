package db

import (
	"tempest-administration-service/pkg/entities"
	"testing"
)

func TestCreateNewUserValidUsernameAndPasswordIsSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	mockDBConn := NewDBConnFromExisting(dbMock)

	err := mockDBConn.CreateUser(entities.User{
		Username:       realTestValueUsername,
		PasswordHashed: realTestValueHashedPass,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestCreateNewUserInvalidUsernameIsNotSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	mockDBConn := NewDBConnFromExisting(dbMock)

	err := mockDBConn.CreateUser(entities.User{
		Username:       "",
		PasswordHashed: realTestValueHashedPass,
	})

	if err == nil {
		t.Errorf("expected failure, username was too short")
	}

}

func TestCreateNewUserInvalidPasswordIsNotSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	mockDBConn := NewDBConnFromExisting(dbMock)

	err := mockDBConn.CreateUser(entities.User{
		Username:       realTestValueUsername,
		PasswordHashed: "",
	})

	if err == nil {
		t.Errorf("expected failure, password was too short")
	}

}
