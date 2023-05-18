package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tempest-administration-service/pkg/config"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartServer(conf *config.Service, router *mux.Router) error {

	handler := cors.AllowAll().Handler(router)

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started on port: %v", conf.Port)

	log.Fatal(srv.ListenAndServe())
	return nil
}
