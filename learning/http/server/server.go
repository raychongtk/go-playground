package server

import (
	"log"
	"net/http"
)

func Start() {
	router := registerRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
