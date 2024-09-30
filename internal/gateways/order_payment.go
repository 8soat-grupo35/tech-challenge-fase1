package gateways

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/repository"
	"gorm.io/gorm"
	"log"
)

type orderPaymentGateway struct {
	orm *gorm.DB
}

func NewOrderPaymentGateway(orm *gorm.DB) repository.OrderPaymentRepository {
	return &orderPaymentGateway{orm: orm}
}

func (o orderPaymentGateway) Create(orderPayment entities.OrderPayment) (*entities.OrderPayment, error) {
	result := o.orm.Create(&orderPayment)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &orderPayment, nil
}

func (o orderPaymentGateway) Update(orderId uint32, orderPayment *entities.OrderPayment) (*entities.OrderPayment, error) {

	orderPaymentToUpdate := entities.OrderPayment{
		PaymentStatusID: orderPayment.PaymentStatusID,
	}

	result := o.orm.Where(entities.OrderPayment{OrderID: orderId}).Updates(&orderPaymentToUpdate)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &orderPaymentToUpdate, nil
}

func (o orderPaymentGateway) GetOneByOrderID(orderID uint32) (orderPayment *entities.OrderPayment, err error) {
	result := o.orm.Preload("PaymentStatus").Where(entities.OrderPayment{OrderID: orderID}).First(&orderPayment)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return orderPayment, nil
}
