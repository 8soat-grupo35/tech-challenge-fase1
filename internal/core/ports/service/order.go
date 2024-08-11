package service

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
)

//go:generate mockgen -source=order.go -destination=../../../../test/core/ports/service/mock/services_mock.go
type OrderService interface {
	GetAll() ([]domain.Order, error)
	Create(order dto.OrderDto) (*domain.Order, error)
}
