package main

import (
	"log"
	"net/http"

	"github.com/cikupin/go-climate-app/routes"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.Router.Handler()
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
