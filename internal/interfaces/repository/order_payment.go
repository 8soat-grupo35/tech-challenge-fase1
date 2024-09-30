package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"

type OrderPaymentRepository interface {
	GetOneByOrderID(orderID uint32) (*entities.OrderPayment, error)
	Create(payment entities.OrderPayment) (*entities.OrderPayment, error)
	Update(orderId uint32, orderPayment *entities.OrderPayment) (*entities.OrderPayment, error)
}
