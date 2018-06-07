package main

import (
	"log"
	"net/http"

	"github.com/cikupin/go-climate-app/routes"
)

func main() {
	r := new(routes.Router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
