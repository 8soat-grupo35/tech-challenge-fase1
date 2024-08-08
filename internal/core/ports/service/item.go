package service

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
)

//go:generate mockgen -source=item.go -destination=../../../../test/core/ports/service/mock/services_mock.go
type ItemService interface {
	GetAll(category string) ([]domain.Item, error)
	Create(item dto.ItemDto) (*domain.Item, error)
	Update(itemId uint32, item dto.ItemDto) (*domain.Item, error)
	Delete(itemId uint32) error
}
