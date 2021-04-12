package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("someone accessed a wrong url", r.URL.RequestURI())
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(200)
	_, err := w.Write([]byte("welcome to my home page"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "something went wrong", 500)
		return
	}

	log.Println("someone accessed home page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
