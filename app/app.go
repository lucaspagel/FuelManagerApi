package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	log.Println("Hello World")

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
