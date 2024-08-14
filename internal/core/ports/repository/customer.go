package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

type CustomerRepository interface {
	GetAll() ([]domain.Customer, error)
	GetOne(domain.Customer) (*domain.Customer, error)
	Create(customer domain.Customer) (*domain.Customer, error)
	Update(customerId uint32, customer domain.Customer) (*domain.Customer, error)
	Delete(customerId uint32) error
}
