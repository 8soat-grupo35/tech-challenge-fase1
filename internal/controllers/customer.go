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

type CustomerController struct {
	dbConnection *gorm.DB
	gateway      repository.CustomerRepository
	useCase      usecase.CustomerUseCase
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	gateway := gateways.NewCustomerGateway(db)
	return &CustomerController{
		dbConnection: db,
		gateway:      gateway,
		useCase:      usecases.NewCustomerUseCase(gateway),
	}
}

func (c *CustomerController) GetAll() ([]entities.Customer, error) {
	return c.useCase.GetAll()
}

func (c *CustomerController) GetByCPF(cpf string) (*entities.Customer, error) {
	return c.useCase.GetByCpf(cpf)
}

func (c *CustomerController) Create(customer dto.CustomerDto) (*entities.Customer, error) {
	return c.useCase.Create(customer)
}

func (c *CustomerController) Update(customerID int, customer dto.CustomerDto) (*entities.Customer, error) {
	return c.useCase.Update(uint32(customerID), customer)
}

func (c *CustomerController) Delete(customerID int) error {
	return c.useCase.Delete(uint32(customerID))
}
