package services

import (
	"errors"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports"
)

type Service struct {
	itemRepository ports.ItemRepository
}

func New(itemRepository ports.ItemRepository) *Service {
	return &Service{
		itemRepository: itemRepository,
	}
}

func (service *Service) GetAll() ([]domain.Item, error) {
	items, err := service.itemRepository.GetAll()

	if err != nil {
		return []domain.Item{}, errors.New("get item from repository has failed")
	}

	return items, nil
}
