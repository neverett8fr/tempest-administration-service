package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	entities "tempest-administration-service/pkg/entitites"

	"github.com/gorilla/mux"
)

func newUserAuth(r *mux.Router) {

	r.HandleFunc("/test/{text}", testHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", createUserHandler).Methods(http.MethodPost)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	text := params["text"]

	body := NewResponse(fmt.Sprintf("test: %v", text))

	writeReponse(w, body)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	bodyIn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}
	userInformation := newUserIn{}
	err = json.Unmarshal(bodyIn, &userInformation)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	user, err := entities.NewUser(userInformation.Username, userInformation.Password)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	err = DBConn.CreateUser(user)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := NewResponse(fmt.Sprintf("user created with username %v", userInformation.Username), err)

	writeReponse(w, body)
}
