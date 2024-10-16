package usecases

import (
	"fmt"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/api/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/interfaces/usecase"
	"log"
)

type orderPaymentUseCase struct {
	orderPaymentRepository       repository.OrderPaymentRepository
	orderPaymentStatusRepository repository.OrderPaymentStatusRepository
}

func NewOrderPaymentUseCase(
	orderPaymentRepository repository.OrderPaymentRepository,
	orderPaymentStatusRepository repository.OrderPaymentStatusRepository,
) usecase.OrderPaymentUseCase {
	return &orderPaymentUseCase{
		orderPaymentRepository:       orderPaymentRepository,
		orderPaymentStatusRepository: orderPaymentStatusRepository,
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

func (o orderPaymentUseCase) UpdateStatus(orderID uint32, status string) (orderPayment *entities.OrderPayment, err error) {
	orderPayment, err = o.orderPaymentRepository.GetOneByOrderID(orderID)

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "get order payment has failed",
		}
	}

	paymentStatus, err := o.orderPaymentStatusRepository.GetByName(status)
	fmt.Println(paymentStatus)
	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "get order payment status has failed",
		}
	}

	err = orderPayment.SetPaymentStatus(paymentStatus.ID)

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	orderPaymentUpdated, err := o.orderPaymentRepository.Update(orderID, orderPayment)

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "update order payment has failed",
		}
	}

	return orderPaymentUpdated, nil
}
