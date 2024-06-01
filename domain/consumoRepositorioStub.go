package domain

import (
	"errors"
	"reflect"
	"time"
)

type ConsumoRepositoryStub struct {
	consumos map[int]Consumo
}

// NewConsumoRepositoryStub Creates dummy consumos
func NewConsumoRepositoryStub() *ConsumoRepositoryStub {
	consumos := map[int]Consumo{
		1: {1, time.Now(), 14.55, Etanol, 1, Veiculo{}},
		2: {2, time.Now(), 200.05, Diesel, 2, Veiculo{}},
		3: {3, time.Now(), 65.00, Gasolina, 1, Veiculo{}},
	}

	return &ConsumoRepositoryStub{consumos}
}

func (s *ConsumoRepositoryStub) GetAll() (map[int]Consumo, error) {
	return s.consumos, nil
}

func (s *ConsumoRepositoryStub) GetById(id int) (Consumo, error) {
	for _, v := range s.consumos {
		if v.Id == id {
			return v, nil
		}
	}
	return Consumo{}, nil
}

func (s *ConsumoRepositoryStub) Create(c Consumo) (Consumo, error) {
	if _, exists := s.consumos[c.Id]; exists {
		return Consumo{}, errors.New("consumo com id ja existente")
	}
	s.consumos[c.Id] = c
	return c, nil
}

func (s *ConsumoRepositoryStub) Update(id int, c Consumo) (Consumo, error) {
	if _, exists := s.consumos[id]; !exists {
		return Consumo{}, errors.New("consumo com id não ja existente")
	}

	newConsumo := Consumo{}

	// Only update non-zero values
	if !c.Data.IsZero() {
		newConsumo.Data = c.Data
	}
	if c.Valor != 0 {
		newConsumo.Valor = c.Valor
	}
	if c.Tipo != 0 {
		newConsumo.Tipo = c.Tipo
	}
	if c.VeiculoId != 0 {
		newConsumo.VeiculoId = c.VeiculoId
	}

	s.consumos[id] = newConsumo

	return c, nil
}

func (s *ConsumoRepositoryStub) PatchUpdate(id int, c Consumo) (Consumo, error) {
	existingConsumo, exists := s.consumos[id]
	if !exists {
		return Consumo{}, errors.New("consumo com id não existente")
	}

	vVal := reflect.ValueOf(c)
	vExisting := reflect.ValueOf(&existingConsumo).Elem()

	// iterating through the fields of c
	for i := 0; i < vVal.NumField(); i++ {
		field := vVal.Field(i)
		if !field.IsZero() {
			vExisting.Field(i).Set(field)
		}
	}

	s.consumos[id] = existingConsumo

	return existingConsumo, nil
}

func (s *ConsumoRepositoryStub) Delete(id int) (Consumo, error) {
	for _, c := range s.consumos {
		if c.Id == id {
			delete(s.consumos, c.Id)

			return c, nil
		}
	}

	return Consumo{}, nil
}
