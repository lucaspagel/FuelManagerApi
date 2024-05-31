package service

import "github.com/lucaspagel/fuel-api/domain"

// Port
type VeiculoService interface {
	GetAllVeiculos() ([]domain.Veiculo, error)
	GetByIdVeiculo(id int) (domain.Veiculo, error)
	CreateVeiculo(v domain.Veiculo) (domain.Veiculo, error)
	UpdateVeiculo(id int, v domain.Veiculo) (domain.Veiculo, error)
	PatchUpdateVeiculo(id int, v domain.Veiculo) (domain.Veiculo, error)
	Delete(id int) (domain.Veiculo, error)
}

type DefaultVeiculoService struct {
	repo domain.VeiculoRepository
}

func NewVeiculoService(repository domain.VeiculoRepository) DefaultVeiculoService {
	return DefaultVeiculoService{repository}
}

func (s DefaultVeiculoService) GetAllVeiculos() ([]domain.Veiculo, error) {
	mapV, _ := s.repo.GetAll()

	veiculos := make([]domain.Veiculo, 0, len(mapV))
	for _, v := range mapV {
		veiculos = append(veiculos, v)
	}
	return veiculos, nil

	// return s.repo.GetAll()
}

func (s DefaultVeiculoService) GetByIdVeiculo(id int) (domain.Veiculo, error) {
	return s.repo.GetById(id)
}

func (s DefaultVeiculoService) CreateVeiculo(v domain.Veiculo) (domain.Veiculo, error) {
	return s.repo.Create(v)
}

func (s DefaultVeiculoService) UpdateVeiculo(id int, v domain.Veiculo) (domain.Veiculo, error) {
	return s.repo.Update(id, v)
}

func (s DefaultVeiculoService) PatchUpdateVeiculo(id int, v domain.Veiculo) (domain.Veiculo, error) {
	return s.repo.PatchUpdate(id, v)
}

func (s DefaultVeiculoService) Delete(id int) (domain.Veiculo, error) {
	return s.repo.Delete(id)
}
