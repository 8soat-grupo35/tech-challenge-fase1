package services

import (
	"errors"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	custom_errors "github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/errors"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) service.OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (service *orderService) GetAll() ([]domain.Order, error) {
	orders, err := service.orderRepository.GetAll()

	if err != nil {
		return []domain.Order{}, &custom_errors.DatabaseError{
			Message: "get order from repository has failed",
		}
	}

	return orders, nil
}

// Create implements ports.OrderService.
func (service *orderService) Create(order dto.OrderDto) (*domain.Order, error) {
	newOrder, err := domain.NewOrder(order)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	orderSaved, err := service.orderRepository.Create(*newOrder)

	if err != nil {
		return nil, errors.New("create order on repository has failed")
	}

	return orderSaved, err
}
