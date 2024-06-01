package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucaspagel/fuel-api/domain"
	"github.com/lucaspagel/fuel-api/service"
	"net/http"
	"strconv"
)

type ConsumoHandler struct {
	service service.ConsumoService
}

func (ch *ConsumoHandler) getAll(w http.ResponseWriter, r *http.Request) {
	veiculos, _ := ch.service.GetAllConsumos()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculos)
}

func (ch *ConsumoHandler) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["consumo_id"]
	id, _ := strconv.Atoi(idStr)

	consumo, _ := ch.service.GetByIdConsumo(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumo)
}

func (ch *ConsumoHandler) create(w http.ResponseWriter, r *http.Request) {
	var request domain.Consumo
	json.NewDecoder(r.Body).Decode(&request)

	consumo, _ := ch.service.CreateConsumo(request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumo)
}

func (ch *ConsumoHandler) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["consumo_id"]
	id, _ := strconv.Atoi(idStr)

	var request domain.Consumo
	json.NewDecoder(r.Body).Decode(&request)

	consumo, _ := ch.service.UpdateConsumo(id, request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumo)
}

func (ch *ConsumoHandler) patchUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["consumo_id"]
	id, _ := strconv.Atoi(idStr)

	var request domain.Consumo
	json.NewDecoder(r.Body).Decode(&request)

	consumo, _ := ch.service.PatchUpdateConsumo(id, request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumo)
}

func (ch *ConsumoHandler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["consumo_id"]
	id, _ := strconv.Atoi(idStr)

	consumo, _ := ch.service.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consumo)
}
