package repository

import "github.com/8soat-grupo35/tech-challenge-fase1/src/entities"

//go:generate mockgen -source=customer.go -destination=../../../test/gateways/mock/customer_mock.go
type CustomerRepository interface {
	GetAll() ([]entities.Customer, error)
	GetOne(entities.Customer) (*entities.Customer, error)
	Create(customer entities.Customer) (*entities.Customer, error)
	Update(customerId uint32, customer entities.Customer) (*entities.Customer, error)
	Delete(customerId uint32) error
}
