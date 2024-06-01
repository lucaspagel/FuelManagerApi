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

	vh := VeiculoHandler{service.NewVeiculoService(domain.NewVeiculoRepositoryStub())}
	ch := ConsumoHandler{service.NewConsumoService(domain.NewConsumoRepositoryStub())}

	router.HandleFunc("/veiculos", vh.getAll).Methods(http.MethodGet)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", vh.getById).Methods(http.MethodGet)
	router.HandleFunc("/veiculos", vh.create).Methods(http.MethodPost)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", vh.update).Methods(http.MethodPut)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", vh.patchUpdate).Methods(http.MethodPatch)
	router.HandleFunc("/veiculos/{veiculo_id:[0-9]+}", vh.delete).Methods(http.MethodDelete)

	router.HandleFunc("/consumos", ch.getAll).Methods(http.MethodGet)
	router.HandleFunc("/consumos/{consumo_id:[0-9]+}", ch.getById).Methods(http.MethodGet)
	router.HandleFunc("/consumos", ch.create).Methods(http.MethodPost)
	router.HandleFunc("/consumos/{consumo_id:[0-9]+}", ch.update).Methods(http.MethodPut)
	router.HandleFunc("/consumos/{consumo_id:[0-9]+}", ch.patchUpdate).Methods(http.MethodPatch)
	router.HandleFunc("/consumos/{consumo_id:[0-9]+}", ch.delete).Methods(http.MethodDelete)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
