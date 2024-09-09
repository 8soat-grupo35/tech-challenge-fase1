package usecase

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
)

type ItemUseCase interface {
	GetAll(category string) ([]entities.Item, error)
	Create(item dto.ItemDto) (*entities.Item, error)
	Update(itemId uint32, item dto.ItemDto) (*entities.Item, error)
	Delete(itemId uint32) error
}