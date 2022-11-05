package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func newAdministrationOperation(r *mux.Router) {

	r.HandleFunc("/token", createToken).Methods(http.MethodPost)
	r.HandleFunc("/token", checkToken).Methods(http.MethodGet)
}

func createToken(w http.ResponseWriter, r *http.Request) {

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

	tok, err := TokenProvider.NewToken(userInformation.Username, userInformation.Password)
	if err != nil {
		body := NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	body := NewResponse(tokenOut{Token: tok}, err)

	writeReponse(w, body)
}

func checkToken(w http.ResponseWriter, r *http.Request) {
	// check headers or vars

	writeReponse(w, nil)
}
