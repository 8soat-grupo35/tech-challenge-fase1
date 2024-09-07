package usecase

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
)

type OrderUseCase interface {
	GetAll() ([]entities.Order, error)
	Create(order dto.OrderDto) (*entities.Order, error)
}
