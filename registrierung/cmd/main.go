package main

import (
	"net/http"

	"github.com/chris21101/trainingcblank/registrierung"
	"github.com/gorilla/mux"
	"training-fellow.de/registrierung"
)

func main() {
	regHandler := &registrierung.RegistrierungsHandler{}
	r := mux.NewRouter()
	r.PathPefix("/").Methods("POST").Handler(regHandler)
	http.ListenAndServe("localhost:8989", r)
}
