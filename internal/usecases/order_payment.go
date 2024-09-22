package usecases

import (
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/api/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/usecase"
)

type orderPaymentUseCase struct {
	orderPaymentRepository repository.OrderPaymentRepository
}

func NewOrderPaymentUseCase(orderPaymentRepository repository.OrderPaymentRepository) usecase.OrderPaymentUseCase {
	return &orderPaymentUseCase{
		orderPaymentRepository: orderPaymentRepository,
	}
}

func (o orderPaymentUseCase) GetPayment(orderID uint32) (*entities.OrderPayment, error) {
	orderPayment, err := o.orderPaymentRepository.GetOneByOrderID(orderID)

	if err != nil {
		return nil, &custom_errors.DatabaseError{
			Message: err.Error(),
		}
	}

	return orderPayment, nil
}

func (o orderPaymentUseCase) Create(order entities.Order) (*entities.OrderPayment, error) {
	newOrderPayment, err := entities.NewOrderPayment(order.ID, entities.PAYMENT_STATUS_WAITING)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	orderPaymentSaved, err := o.orderPaymentRepository.Create(*newOrderPayment)

	if err != nil {
		return nil, &custom_errors.DatabaseError{
			Message: err.Error(),
		}
	}

	return orderPaymentSaved, nil
}
