package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateNewUserIsSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	router := newMockRouter(t, dbMock)

	newRequest := newUserIn{
		Username: realTestValueUsername,
		Password: realTestValuePass,
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(newRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	path := "/user"
	req, err := http.NewRequest(http.MethodPost, path, &buf)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error, expected %v, got %v", http.StatusOK, rr.Code)
	}
}

func TestCreateNewUserInvalidUsernameIsNotSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	router := newMockRouter(t, dbMock)

	newRequest := newUserIn{
		Username: "",
		Password: realTestValuePass,
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(newRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	path := "/user"
	req, err := http.NewRequest(http.MethodPost, path, &buf)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code == http.StatusOK {
		t.Errorf("error, expected %v, got %v", http.StatusOK, rr.Code)
	}
}

func TestCreateNewUserInvalidPasswordIsNotSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	router := newMockRouter(t, dbMock)

	newRequest := newUserIn{
		Username: realTestValueUsername,
		Password: "",
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(newRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	path := "/user"
	req, err := http.NewRequest(http.MethodPost, path, &buf)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code == http.StatusOK {
		t.Errorf("error, expected %v, got %v", http.StatusOK, rr.Code)
	}
}
func TestCreateNewUserNoBodyIsNotSuccessful(t *testing.T) {

	dbMock, mock := newMockDB(t)
	newMockAddUser(t, mock)
	defer dbMock.Close()

	router := newMockRouter(t, dbMock)

	newRequest := newUserIn{}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(newRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	path := "/user"
	req, err := http.NewRequest(http.MethodPost, path, &buf)
	if err != nil {
		t.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code == http.StatusOK {
		t.Errorf("error, expected %v, got %v", http.StatusOK, rr.Code)
	}
}
