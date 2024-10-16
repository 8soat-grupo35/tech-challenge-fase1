package usecase

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"

type OrderPaymentUseCase interface {
	Create(order entities.Order) (*entities.OrderPayment, error)
	GetPayment(orderID uint32) (*entities.OrderPayment, error)
	UpdateStatus(orderID uint32, status string) (*entities.OrderPayment, error)
}
