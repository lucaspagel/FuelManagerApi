package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucaspagel/fuel-api/domain"
	"github.com/lucaspagel/fuel-api/service"
	"net/http"
	"strconv"
)

type VeiculoHandler struct {
	service service.VeiculoService
}

func (ch *VeiculoHandler) getAll(w http.ResponseWriter, r *http.Request) {
	veiculos, _ := ch.service.GetAllVeiculos()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculos)
}

func (ch *VeiculoHandler) getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["veiculo_id"]
	id, _ := strconv.Atoi(idStr)

	veiculo, _ := ch.service.GetByIdVeiculo(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}

func (ch *VeiculoHandler) create(w http.ResponseWriter, r *http.Request) {
	var request domain.Veiculo
	json.NewDecoder(r.Body).Decode(&request)

	veiculo, _ := ch.service.CreateVeiculo(request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}

func (ch *VeiculoHandler) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["veiculo_id"]
	id, _ := strconv.Atoi(idStr)

	var request domain.Veiculo
	json.NewDecoder(r.Body).Decode(&request)

	veiculo, _ := ch.service.UpdateVeiculo(id, request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}

func (ch *VeiculoHandler) patchUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["veiculo_id"]
	id, _ := strconv.Atoi(idStr)

	var request domain.Veiculo
	json.NewDecoder(r.Body).Decode(&request)

	veiculo, _ := ch.service.PatchUpdateVeiculo(id, request)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}

func (ch *VeiculoHandler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["veiculo_id"]
	id, _ := strconv.Atoi(idStr)

	veiculo, _ := ch.service.Delete(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(veiculo)
}
