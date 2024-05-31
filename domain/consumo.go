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
