package usecase

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
)

type OrderUseCase interface {
	GetAll() ([]entities.Order, error)
	Create(order dto.OrderDto) (*entities.Order, error)
	UpdateStatus(id uint32, status string) (*entities.Order, error)
}
