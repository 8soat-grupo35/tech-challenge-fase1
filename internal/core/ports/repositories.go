package ports

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type ItemRepository interface {
	GetAll() ([]domain.Item, error)
}
