package main

import (
	"database/sql"
	"log"
	"tempest-administration-service/cmd"
	"tempest-administration-service/pkg/application/service"
	"tempest-administration-service/pkg/config"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes(conn *sql.DB, conf config.Config) *mux.Router {
	r := mux.NewRouter()
	service.NewServiceRoutes(r, conn, conf)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	serviceDB, err := cmd.OpenDB(&conf.DB)
	if err != nil {
		log.Fatalf("error starting db, err %v", err)
		return
	}
	defer serviceDB.Close()
	log.Println("connection to DB setup")

	err = cmd.MigrateDB(serviceDB, conf.DB.Driver)
	if err != nil {
		log.Fatalf("error running DB migrations, %v", err)
		return
	}
	log.Println("DB migrations ran")

	router := getRoutes(serviceDB, *conf)
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}
