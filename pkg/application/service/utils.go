package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	autho "tempest-administration-service/pkg/infra/auth"
	"tempest-administration-service/pkg/infra/db"

	"github.com/gorilla/mux"
)

var (
	DBConn        *db.DBConn
	TokenProvider autho.TokenProvider
)

type Response struct {
	Data   interface{} `json:"data"`
	Errors []error     `json:"errors"`
}

func NewServiceRoutes(r *mux.Router, conn *sql.DB) {
	DBConn = db.NewDBConnFromExisting(conn)
	TokenProvider = autho.InitialiseTokenProvider(DBConn)

	// newAdministrationInformation(r)
	// newAdministrationOperation(r)
	newUserAuth(r)
	newAdministrationOperation(r)
}

func NewResponse(data interface{}, err ...error) *Response {

	return &Response{
		Data:   data,
		Errors: err,
	}
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
		return
	}
}
