package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"

type OrderPaymentStatusRepository interface {
	GetByName(string) (*entities.PaymentStatus, error)
}
