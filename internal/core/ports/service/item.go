package service

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

//go:generate mockgen -source=item.go -destination=../../../../test/core/ports/service/mock/services_mock.go
type ItemService interface {
	GetAll(domain.Item) ([]domain.Item, error)
	Create(domain.Item) (*domain.Item, error)
	Update(itemId uint32, item domain.Item) (*domain.Item, error)
	Delete(itemId uint32) error
}
