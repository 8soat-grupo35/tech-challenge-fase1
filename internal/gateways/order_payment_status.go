package gateways

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/repository"
	"gorm.io/gorm"
)

type orderPaymentStatusGateway struct {
	orm *gorm.DB
}

func NewOrderPaymentStatusGateway(orm *gorm.DB) repository.OrderPaymentStatusRepository {
	return &orderPaymentStatusGateway{orm: orm}
}

func (o orderPaymentStatusGateway) GetByName(s string) (paymentStatus *entities.PaymentStatus, err error) {
	result := o.orm.Where("name = ?", s).First(&paymentStatus)

	if result.Error != nil {
		return nil, result.Error
	}

	return paymentStatus, nil
}
