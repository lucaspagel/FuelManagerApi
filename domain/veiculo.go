package domain

type Veiculo struct {
	Id            int
	Marca         string
	Modelo        string
	Placa         string
	AnoFabricacao int
	AnoModelo     int

	Consumos []Consumo
}

// Port
type VeiculoRepository interface {
	GetAll() (map[int]Veiculo, error)
	GetById(id int) (Veiculo, error)
	Create(v Veiculo) (Veiculo, error)
	Update(id int, v Veiculo) (Veiculo, error)
	PatchUpdate(id int, v Veiculo) (Veiculo, error)
	Delete(id int) (Veiculo, error)
}
