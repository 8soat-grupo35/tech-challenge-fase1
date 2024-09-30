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
	orderPaymentStatusGateway := gateways.NewOrderPaymentStatusGateway(db)
	return &OrderController{
		orderUseCase: usecases.NewOrderUseCase(orderGateway),
		orderPaymentUseCase: usecases.NewOrderPaymentUseCase(
			orderPaymentGateway,
			orderPaymentStatusGateway,
		),
	}
}

func (o *OrderController) GetAll() ([]entities.Order, error) {

	return o.orderUseCase.GetAll()
}

func (o *OrderController) Checkout(orderDto dto.OrderDto) (*presenters.OrderPresenter, error) {
	order, err := o.orderUseCase.Create(orderDto)

	if err != nil {
		return nil, err
	}

	_, err = o.orderPaymentUseCase.Create(*order)

	if err != nil {
		return nil, err
	}

	return &presenters.OrderPresenter{Id: order.ID}, nil
}

func (o *OrderController) GetPaymentStatus(orderID uint32) (*presenters.OrderPaymentStatusPresenter, error) {
	orderPayment, err := o.orderPaymentUseCase.GetPayment(orderID)
	if err != nil {
		return nil, &custom_errors.NotFoundError{
			Message: err.Error(),
		}
	}

	return &presenters.OrderPaymentStatusPresenter{
		PaymentStatus: orderPayment.PaymentStatus.Name,
	}, nil
}

func (o *OrderController) UpdatePaymentStatus(orderID uint32, status string) error {
	_, err := o.orderPaymentUseCase.UpdateStatus(orderID, status)

	return err
}
