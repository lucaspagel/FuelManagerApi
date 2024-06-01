package domain

import (
	"time"
)

// Define the TipoCombustivel enum
type TipoCombustivel int

const (
	Diesel TipoCombustivel = iota
	Etanol
	Gasolina
)

type Consumo struct {
	Id    int
	Data  time.Time
	Valor float64
	Tipo  TipoCombustivel

	VeiculoId int

	Veiculo Veiculo
}

// Port
type ConsumoRepository interface {
	GetAll() (map[int]Consumo, error)
	GetById(id int) (Consumo, error)
	Create(v Consumo) (Consumo, error)
	Update(id int, v Consumo) (Consumo, error)
	PatchUpdate(id int, v Consumo) (Consumo, error)
	Delete(id int) (Consumo, error)
}
