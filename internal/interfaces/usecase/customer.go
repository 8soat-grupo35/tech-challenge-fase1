package usecase

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/src/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
)

type CustomerUseCase interface {
	GetAll() ([]entities.Customer, error)
	Create(dto.CustomerDto) (*entities.Customer, error)
	GetByCpf(cpf string) (*entities.Customer, error)
	Update(customerId uint32, customer dto.CustomerDto) (*entities.Customer, error)
	Delete(customerId uint32) error
}
