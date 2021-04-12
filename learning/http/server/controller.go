package server

import (
	"encoding/json"
	"http/server/api"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("welcome to my home page"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "something went wrong", 500)
		return
	}

	log.Println("someone accessed home page")
}

func login(w http.ResponseWriter, r *http.Request) {
	var loginRequest api.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "invalid params", 400)
		return
	}

	var loginResponse api.LoginResponse
	if loginRequest.Username == "ray" && loginRequest.Password == "123456" {
		w.WriteHeader(200)
		loginResponse.Success = true
	} else {
		w.WriteHeader(401)
		var errorMessage string
		errorMessage = "username or password wrong"
		loginResponse.Success = false
		loginResponse.ErrorMessage = &errorMessage
	}
	json.NewEncoder(w).Encode(&loginResponse)
}
