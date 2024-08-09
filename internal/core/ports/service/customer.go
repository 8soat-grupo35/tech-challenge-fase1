package service

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type CustomerService interface {
	GetAll() ([]domain.Customer, error)
	Create(domain.Customer) (*domain.Customer, error)
	GetByCpf(cpf string) (*domain.Customer, error)
	Update(customerId uint32, customer domain.Customer) (*domain.Customer, error)
	Delete(customerId uint32) error
}
