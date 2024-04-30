package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = "8083"

func main() {
	app := Config{}
	fmt.Println("Starting mail service on port" + webPort)

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
