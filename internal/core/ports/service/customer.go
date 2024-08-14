package service

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
)

type CustomerService interface {
	GetAll() ([]domain.Customer, error)
	Create(dto.CustomerDto) (*domain.Customer, error)
	GetByCpf(cpf string) (*domain.Customer, error)
	Update(customerId uint32, customer dto.CustomerDto) (*domain.Customer, error)
	Delete(customerId uint32) error
}
