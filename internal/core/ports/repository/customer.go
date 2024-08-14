package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

//go:generate mockgen -source=customer.go -destination=../../../../test/core/ports/repository/mock/customer_mock.go
type CustomerRepository interface {
	GetAll() ([]domain.Customer, error)
	GetOne(domain.Customer) (*domain.Customer, error)
	Create(customer domain.Customer) (*domain.Customer, error)
	Update(customerId uint32, customer domain.Customer) (*domain.Customer, error)
	Delete(customerId uint32) error
}
