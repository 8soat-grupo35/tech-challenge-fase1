package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

//go:generate mockgen -source=order.go -destination=../../../../test/core/ports/repository/mock/order_mock.go
type OrderRepository interface {
	GetAll() ([]domain.Order, error)
	Create(order domain.Order) (*domain.Order, error)
}
