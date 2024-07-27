package ports

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type ItemService interface {
	GetAll(domain.Item) ([]domain.Item, error)
	Create(domain.Item) (*domain.Item, error)
	Update(itemId uint32, item domain.Item) (*domain.Item, error)
	Delete(itemId uint32) error
}
