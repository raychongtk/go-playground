package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func registerRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/login", login).Methods(http.MethodPut)
	return router
}
