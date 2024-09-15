package controllers

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/api/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/gateways"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/usecase"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/presenters"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/usecases"
	"gorm.io/gorm"
)

type OrderController struct {
	orderUseCase        usecase.OrderUseCase
	orderPaymentUseCase usecase.OrderPaymentUseCase
}

func NewOrderController(db *gorm.DB) *OrderController {
	orderGateway := gateways.NewOrderGateway(db)
	orderPaymentGateway := gateways.NewOrderPaymentGateway(db)
	return &OrderController{
		orderUseCase:        usecases.NewOrderUseCase(orderGateway),
		orderPaymentUseCase: usecases.NewOrderPaymentUseCase(orderPaymentGateway),
	}
}

func (o *OrderController) GetAll() ([]entities.Order, error) {

	return o.orderUseCase.GetAll()
}

func (o *OrderController) Checkout(orderDto dto.OrderDto) (*entities.Order, error) {
	order, err := o.orderUseCase.Create(orderDto)

	if err != nil {
		return nil, err
	}

	_, err = o.orderPaymentUseCase.Create(*order)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderController) GetPaymentStatus(orderID uint32) (*presenters.OrderPaymentStatusPresenter, error) {
	orderPayment, err := o.orderPaymentUseCase.GetPayment(orderID)
	if err != nil {
		return nil, &custom_errors.NotFoundError{
			Message: "order not found",
		}
	}

	return &presenters.OrderPaymentStatusPresenter{
		PaymentStatus: orderPayment.PaymentStatus.Name,
	}, nil
}
