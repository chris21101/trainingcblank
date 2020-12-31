package main

import (
	"net/http"

	"github.com/chris21101/trainingcblank/registrierung"
	"github.com/gorilla/mux"
)

func main() {
	regHandler := &registrierung.RegistrierungsHandler{}
	r := mux.NewRouter()
	r.Handle("/", regHandler)
	r.Methods("POST")
	http.ListenAndServe("localhost:8989", r)
}
