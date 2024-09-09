package controllers

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/gateways"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/interfaces/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/interfaces/usecase"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/usecases"
	"gorm.io/gorm"
)

type OrderController struct {
	dbConnection *gorm.DB
	gateway      repository.OrderRepository
	useCase      usecase.OrderUseCase
}

func NewOrderController(db *gorm.DB) *OrderController {
	gateway := gateways.NewOrderGateway(db)
	return &OrderController{
		dbConnection: db,
		gateway:      gateway,
		useCase:      usecases.NewOrderUseCase(gateway),
	}
}

func (o *OrderController) GetAll() ([]entities.Order, error) {
	return o.useCase.GetAll()
}

func (o *OrderController) Checkout(orderDto dto.OrderDto) (*entities.Order, error) {
	return o.useCase.Create(orderDto)
}