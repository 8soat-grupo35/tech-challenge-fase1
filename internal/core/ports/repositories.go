package ports

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type ItemRepository interface {
	GetAll() ([]domain.Item, error)
	GetOne(domain.Item) (*domain.Item, error)
	Create(item domain.Item) (*domain.Item, error)
	Update(itemId uint32, item domain.Item) (*domain.Item, error)
	Delete(itemId uint32) error
}
