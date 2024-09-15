package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"

//go:generate mockgen -source=item.go -destination=../../../test/gateways/mock/item_mock.go
type ItemRepository interface {
	GetAll(entities.Item) ([]entities.Item, error)
	GetOne(entities.Item) (*entities.Item, error)
	Create(item entities.Item) (*entities.Item, error)
	Update(itemId uint32, item entities.Item) (*entities.Item, error)
	Delete(itemId uint32) error
}
