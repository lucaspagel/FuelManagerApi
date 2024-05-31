package domain

import (
	"errors"
	"reflect"
)

type VeiculoRepositoryStub struct {
	veiculos map[int]Veiculo
}

// NewVeiculoRepositoryStub Creates dummy veiculos
func NewVeiculoRepositoryStub() *VeiculoRepositoryStub {
	veiculos := map[int]Veiculo{
		1: {1, "Chevrolet", "Sonic", "BRA2E19", 2014, 2014, nil},
		2: {2, "Chevrolet", "Zafira", "OTM2022", 2009, 2010, nil},
	}

	return &VeiculoRepositoryStub{veiculos}
}

func (s *VeiculoRepositoryStub) GetAll() (map[int]Veiculo, error) {
	return s.veiculos, nil
}

func (s *VeiculoRepositoryStub) GetById(id int) (Veiculo, error) {
	for _, v := range s.veiculos {
		if v.Id == id {
			return v, nil
		}
	}
	return Veiculo{}, nil
}

func (s *VeiculoRepositoryStub) Create(v Veiculo) (Veiculo, error) {
	if _, exists := s.veiculos[v.Id]; exists {
		return Veiculo{}, errors.New("veiculo com id ja existente")
	}
	s.veiculos[v.Id] = v
	return v, nil
}

func (s *VeiculoRepositoryStub) Update(id int, v Veiculo) (Veiculo, error) {
	if _, exists := s.veiculos[id]; !exists {
		return Veiculo{}, errors.New("veiculo com id não ja existente")
	}

	newVeiculo := Veiculo{}

	// Only update non-zero values
	if v.Marca != "" {
		newVeiculo.Marca = v.Marca
	}
	if v.Modelo != "" {
		newVeiculo.Modelo = v.Modelo
	}
	if v.Placa != "" {
		newVeiculo.Placa = v.Placa
	}
	if v.AnoFabricacao != 0 {
		newVeiculo.AnoFabricacao = v.AnoFabricacao
	}
	if v.AnoModelo != 0 {
		newVeiculo.AnoModelo = v.AnoModelo
	}
	if v.Consumos != nil {
		newVeiculo.Consumos = v.Consumos
	}

	s.veiculos[id] = newVeiculo

	return v, nil
}

func (s *VeiculoRepositoryStub) PatchUpdate(id int, v Veiculo /*updates map[string]interface{}*/) (Veiculo, error) {
	existingVeiculo, exists := s.veiculos[id]
	if !exists {
		return Veiculo{}, errors.New("veiculo com id não existente")
	}

	vVal := reflect.ValueOf(v)
	vExisting := reflect.ValueOf(&existingVeiculo).Elem()

	// iterating through the fields of v
	for i := 0; i < vVal.NumField(); i++ {
		field := vVal.Field(i)
		if !field.IsZero() {
			vExisting.Field(i).Set(field)
		}
	}

	s.veiculos[id] = existingVeiculo

	return existingVeiculo, nil
}

func (s *VeiculoRepositoryStub) Delete(id int) (Veiculo, error) {
	for _, v := range s.veiculos {
		if v.Id == id {
			delete(s.veiculos, v.Id)

			return v, nil
		}
	}

	return Veiculo{}, nil
}
