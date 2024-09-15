package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type OrderPayment struct {
	ID              uint32 `gorm:"primary_key;auto_increment"`
	OrderID         uint32
	PaymentStatusID uint32
	PaymentStatus   PaymentStatus
	gorm.Model
}

func NewOrderPayment(orderID uint32, paymentStatusID uint32) (*OrderPayment, error) {
	orderPayment := &OrderPayment{
		OrderID:         orderID,
		PaymentStatusID: paymentStatusID,
	}

	if err := orderPayment.Validate(); err != nil {
		return nil, err
	}

	return orderPayment, nil
}

func (orderPayment OrderPayment) Validate() error {
	return validation.ValidateStruct(
		&orderPayment,
		validation.Field(
			&orderPayment.OrderID,
			validation.Required,
		),
		validation.Field(
			&orderPayment.PaymentStatusID,
			validation.Required,
			validation.In(uint32(1), uint32(2), uint32(3)),
		),
	)
}

type PaymentStatus struct {
	ID   uint32 `gorm:"primary_key;auto_increment"`
	Name string `gorm:"size:255;not null"`
	gorm.Model
}
