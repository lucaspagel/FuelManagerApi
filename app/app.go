package app

import (
	"github.com/gorilla/mux"
	"github.com/lucaspagel/fuel-api/domain"
	"github.com/lucaspagel/fuel-api/service"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := VeiculoHandler{service.NewVeiculoService(domain.NewVeiculoRepositoryStub())}

	router.HandleFunc("/veiculos", ch.getAll).Methods(http.MethodGet)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", ch.getById).Methods(http.MethodGet)
	router.HandleFunc("/veiculos", ch.create).Methods(http.MethodPost)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", ch.update).Methods(http.MethodPut)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", ch.patchUpdate).Methods(http.MethodPatch)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", ch.delete).Methods(http.MethodDelete)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
