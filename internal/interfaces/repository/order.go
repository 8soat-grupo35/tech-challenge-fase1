package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"

//go:generate mockgen -source=order.go -destination=../../../test/gateways/mock/order_mock.go
type OrderRepository interface {
	GetAll() ([]entities.Order, error)
	GetById(id uint32) (*entities.Order, error)
	Create(order entities.Order) (*entities.Order, error)
	Update(id uint32, order entities.Order) (*entities.Order, error)
}
