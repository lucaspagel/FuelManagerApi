package service

import "github.com/lucaspagel/fuel-api/domain"

// Port
type ConsumoService interface {
	GetAllConsumos() ([]domain.Consumo, error)
	GetByIdConsumo(id int) (domain.Consumo, error)
	CreateConsumo(v domain.Consumo) (domain.Consumo, error)
	UpdateConsumo(id int, v domain.Consumo) (domain.Consumo, error)
	PatchUpdateConsumo(id int, v domain.Consumo) (domain.Consumo, error)
	Delete(id int) (domain.Consumo, error)
}

type DefaultConsumoService struct {
	repo domain.ConsumoRepository
}

func NewConsumoService(repository domain.ConsumoRepository) DefaultConsumoService {
	return DefaultConsumoService{repository}
}

func (s DefaultConsumoService) GetAllConsumos() ([]domain.Consumo, error) {
	mapC, _ := s.repo.GetAll()

	consumos := make([]domain.Consumo, 0, len(mapC))
	for _, v := range mapC {
		consumos = append(consumos, v)
	}
	return consumos, nil
}

func (s DefaultConsumoService) GetByIdConsumo(id int) (domain.Consumo, error) {
	return s.repo.GetById(id)
}

func (s DefaultConsumoService) CreateConsumo(v domain.Consumo) (domain.Consumo, error) {
	return s.repo.Create(v)
}

func (s DefaultConsumoService) UpdateConsumo(id int, v domain.Consumo) (domain.Consumo, error) {
	return s.repo.Update(id, v)
}

func (s DefaultConsumoService) PatchUpdateConsumo(id int, v domain.Consumo) (domain.Consumo, error) {
	return s.repo.PatchUpdate(id, v)
}

func (s DefaultConsumoService) Delete(id int) (domain.Consumo, error) {
	return s.repo.Delete(id)
}
