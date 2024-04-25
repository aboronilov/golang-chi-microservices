package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
)

const webPort = "8081"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service on port ", webPort)
	// ToDo connect to DB

	app := Config{}

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
